package cat

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sync"
	"time"

	"github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/pkg/trace/schema"
	"github.com/tendermint/tendermint/types"
)

// FetchTxsFromKeys is called upon by consensus upon receiving a complete compact block.
// The method iterates through the keys in the compact block. For the transactions it
// already has it adds them to a list. For the transactions that are missing it uses a
// block request to track and retrieve them. Once all transactions are retrieved, it returns
// the complete set to the consensus engine. This can be called multiple times sequentially
// with the  same blockID and is thread safe
func (memR *Reactor) FetchTxsFromKeys(ctx context.Context, blockID []byte, compactData [][]byte) ([][]byte, error) {
	if request, ok := memR.blockFetcher.GetRequest(blockID); ok {
		memR.Logger.Debug("tracking existing request for block transactions")
		// we already have a request for this block
		return request.WaitForBlock(ctx)
	}

	txs := make([][]byte, len(compactData))
	missingKeys := make(map[int]types.TxKey, len(compactData))

	// iterate through the keys to know what transactions we have and what are missing
	for i, key := range compactData {
		txKey, err := types.TxKeyFromBytes(key)
		if err != nil {
			return nil, fmt.Errorf("incorrect compact blocks format: %w", err)
		}
		wtx := memR.mempool.store.get(txKey)
		if wtx != nil {
			txs[i] = wtx.tx
			memR.mempool.store.markAsUnevictable(txKey)
		} else {
			missingKeys[i] = txKey
		}
	}
	memR.Logger.Info("fetching transactions from peers", "numTxs", len(txs), "numMissing", len(missingKeys))
	memR.mempool.metrics.MissingTxs.Add(float64(len(missingKeys)))

	// Check if we got lucky and already had all the transactions.
	if len(missingKeys) == 0 {
		schema.WriteMempoolRecoveryStats(
			memR.traceClient,
			0,
			0,
			len(compactData),
			0,
			nil,
		)
		return txs, nil
	}
	initialNumMissing := len(missingKeys)
	missingTxs := make([]string, len(missingKeys))
	for i, tx := range missingKeys {
		missingTxs[i] = bytes.HexBytes(tx[:]).String()
	}

	// setup a request for this block and begin to track and retrieve all missing transactions
	request := memR.blockFetcher.newRequest(
		blockID,
		memR.mempool.Height(),
		missingKeys,
		txs,
	)

	defer func(missingTxs []string) {
		timeTaken := request.TimeTaken()

		schema.WriteMempoolRecoveryStats(
			memR.traceClient,
			initialNumMissing,
			initialNumMissing-len(request.missingKeys),
			len(compactData),
			timeTaken,
			missingTxs,
		)

		memR.Logger.Info("fetched txs", "timeTaken", timeTaken, "numRetrieved", initialNumMissing-len(request.missingKeys), "numMissing", len(request.missingKeys))
		memR.mempool.metrics.RecoveryRate.Observe(float64(initialNumMissing-len(request.missingKeys)) / float64(initialNumMissing))
	}(missingTxs)

	// request the missing transactions if we haven't already
	for _, key := range missingKeys {
		numHaveTx := len(memR.mempool.seenByPeersSet.Get(key))
		memR.Logger.Info("searching for missing tx", "numHaveTx", numHaveTx)
		memR.findNewPeerToRequestTx(key, 5)
	}

	// Wait for the reactor to retrieve and post all transactions.
	return request.WaitForBlock(ctx)
}

func (memR *Reactor) FetchTxsFromKeysSync(compactData [][]byte) ([][]byte, error) {
	txs := make([][]byte, len(compactData))
	missingKeys := make(map[int]types.TxKey, len(compactData))

	// iterate through the keys to know what transactions we have and what are missing
	for i, key := range compactData {
		txKey, err := types.TxKeyFromBytes(key)
		if err != nil {
			return nil, fmt.Errorf("incorrect compact blocks format: %w", err)
		}
		wtx := memR.mempool.store.get(txKey)
		if wtx != nil {
			txs[i] = wtx.tx
			memR.mempool.store.markAsUnevictable(txKey)
		} else {
			missingKeys[i] = txKey
		}
	}
	memR.Logger.Info("fetching transactions from peers", "numTxs", len(txs), "numMissing", len(missingKeys))
	memR.mempool.metrics.MissingTxs.Add(float64(len(missingKeys)))

	// Check if we got lucky and already had all the transactions.
	if len(missingKeys) == 0 {
		schema.WriteMempoolRecoveryStats(
			memR.traceClient,
			0,
			0,
			len(compactData),
			0,
			nil,
		)
		return txs, nil
	}

	missingTxs := make([]string, len(missingKeys))
	for i, tx := range missingKeys {
		missingTxs[i] = bytes.HexBytes(tx[:]).String()
	}

	schema.WriteMempoolRecoveryStats(
		memR.traceClient,
		len(missingKeys),
		len(missingKeys),
		len(compactData),
		0,
		missingTxs,
	)

	// Wait for the reactor to retrieve and post all transactions.
	return nil, fmt.Errorf("missing transaction: %d", len(missingTxs))
}

// FetchKeysFromTxs is in many ways the opposite method. It takes a full block generated by the application
// and reduces it to the set of keys that need to be gossiped from one mempool to another nodes mempool
// in order to recreate the full block.
func (memR *Reactor) FetchKeysFromTxs(ctx context.Context, txs [][]byte) ([][]byte, error) {
	keys := make([][]byte, len(txs))
	for idx, tx := range txs {
		// check if the context has been cancelled
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		key := [32]byte{}
		blobTx, isBlobTx := types.UnmarshalBlobTx(tx)
		if isBlobTx {
			key = sha256.Sum256(blobTx.Tx)
		} else {
			key = sha256.Sum256(tx)
		}
		keys[idx] = key[:]
		has := memR.mempool.store.has(key)
		if !has {
			// If the mempool provided the initial transactions yet received from
			// consensus a transaction it doesn't recognize, this implies that
			// either a tx was mutated or was added by the application. In either
			// case, it is likely no other mempool has this transaction so we
			// preemptively broadcast it to all other peers
			//
			// We don't set the priority, gasWanted or sender fields because we
			// don't know them.
			wtx := newWrappedTx(tx, key, memR.mempool.Height(), 0, 0, "", isBlobTx)
			wtx.evictable = false
			memR.broadcastNewTx(wtx)
			// For safety we also store this transaction in the mempool (ignoring
			// all size limits) so that we can retrieve it later if needed. Note
			// as we're broadcasting it to all peers, we should not receive a `WantTx`
			// unless it gets rejected by the application in CheckTx.
			//
			// Consensus will have an in memory copy of the entire block which includes
			// this transaction so it should not need it.
			memR.mempool.store.set(wtx)
		}
	}

	// return the keys back to the consensus engine
	return keys, nil
}

type blockFetcher struct {
	// mutex to manage concurrent calls to different parts
	mtx sync.Mutex
	// requests are a map of all processing block requests
	// by blockID.
	requests map[string]*blockRequest
}

// newBlockFetcher returns a new blockFetcher for managing block requests
func newBlockFetcher() *blockFetcher {
	return &blockFetcher{
		requests: make(map[string]*blockRequest),
	}
}

func (bf *blockFetcher) GetRequest(blockID []byte) (*blockRequest, bool) {
	bf.mtx.Lock()
	defer bf.mtx.Unlock()
	request, ok := bf.requests[string(blockID)]
	return request, ok
}

// NewRequest creates a new block request and returns it.
// If a request already exists it returns that instead
func (bf *blockFetcher) newRequest(
	blockID []byte,
	height int64,
	missingKeys map[int]types.TxKey,
	txs [][]byte,
) *blockRequest {
	bf.mtx.Lock()
	defer bf.mtx.Unlock()
	if request, ok := bf.requests[string(blockID)]; ok {
		return request
	}
	request := newBlockRequest(height, missingKeys, txs)
	bf.requests[string(blockID)] = request
	bf.pruneOldRequests(height)
	return request
}

// TryAddMissingTx loops through all current requests and tries to add
// the given transaction (if it is missing).
func (bf *blockFetcher) TryAddMissingTx(key types.TxKey, tx []byte) {
	bf.mtx.Lock()
	defer bf.mtx.Unlock()
	for _, request := range bf.requests {
		request.TryAddMissingTx(key, tx)
	}
}

// IsMissingTx checks if a given transaction is missing from any of the block requests.
func (bf *blockFetcher) IsMissingTx(key types.TxKey) bool {
	bf.mtx.Lock()
	defer bf.mtx.Unlock()
	for _, request := range bf.requests {
		_, ok := request.missingKeys[key.String()]
		if ok {
			return true
		}
	}
	return false
}

// PruneOldRequests removes any requests that are older than the given height.
func (bf *blockFetcher) pruneOldRequests(height int64) {
	for blockID, request := range bf.requests {
		if request.height < height {
			delete(bf.requests, blockID)
		}
	}
}

// blockRequests handle the lifecycle of individual block requests.
type blockRequest struct {
	// immutable fields
	height int64
	doneCh chan struct{}

	mtx sync.Mutex
	// track the remaining keys that are missing
	missingKeysByIndex map[int]types.TxKey
	missingKeys        map[string]int
	// the txs in the block
	txs [][]byte

	// used for metrics
	startTime time.Time
	endTime   time.Time
}

func newBlockRequest(
	height int64,
	missingKeys map[int]types.TxKey,
	txs [][]byte,
) *blockRequest {
	mk := make(map[string]int, len(missingKeys))
	for i, key := range missingKeys {
		mk[key.String()] = i
	}
	return &blockRequest{
		height:             height,
		missingKeysByIndex: missingKeys,
		missingKeys:        mk,
		txs:                txs,
		doneCh:             make(chan struct{}),
		startTime:          time.Now().UTC(),
	}
}

// WaitForBlock is a blocking call that waits for the block to be fetched and completed.
// It can be called concurrently. If the block was already fetched it returns immediately.
func (br *blockRequest) WaitForBlock(ctx context.Context) ([][]byte, error) {
	defer br.setEndTime()
	if br.IsDone() {
		return br.txs, nil
	}

	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-br.doneCh:
			return br.txs, nil
		}
	}
}

// TryAddMissingTx checks if a given transactions was missing and if so
// adds it to the block request.
func (br *blockRequest) TryAddMissingTx(key types.TxKey, tx []byte) bool {
	br.mtx.Lock()
	defer br.mtx.Unlock()
	if index, ok := br.missingKeys[key.String()]; ok {
		delete(br.missingKeys, key.String())
		delete(br.missingKeysByIndex, index)
		br.txs[index] = tx
		// check if there is any more transactions remaining
		if len(br.missingKeys) == 0 {
			// Yaay! We're done!
			close(br.doneCh)
		}
		return true
	}
	return false
}

// IsDone returns whether all transactions in the block have been received.
// This is done by measuring the amount of missing keys.
func (br *blockRequest) IsDone() bool {
	br.mtx.Lock()
	defer br.mtx.Unlock()
	return len(br.missingKeys) == 0
}

func (br *blockRequest) TimeTaken() uint64 {
	if br.endTime.IsZero() {
		return 0
	}
	return uint64(br.endTime.Sub(br.startTime).Milliseconds())
}

func (br *blockRequest) setEndTime() {
	br.mtx.Lock()
	defer br.mtx.Unlock()
	br.endTime = time.Now().UTC()
}
