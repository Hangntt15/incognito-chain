package blockchain

import (
	"time"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/metadata"
)

type ShardToBeaconPool interface {
	RemoveBlock(map[byte]uint64)
	//GetFinalBlock() map[byte][]ShardToBeaconBlock
	AddShardToBeaconBlock(*ShardToBeaconBlock) (uint64, uint64, error)
	//ValidateShardToBeaconBlock(ShardToBeaconBlock) error
	GetValidBlockHash() map[byte][]common.Hash
	GetValidBlock(map[byte]uint64) map[byte][]*ShardToBeaconBlock
	GetValidBlockHeight() map[byte][]uint64
	GetLatestValidPendingBlockHeight() map[byte]uint64
	GetBlockByHeight(shardID byte, height uint64) *ShardToBeaconBlock
	SetShardState(map[byte]uint64)
	GetAllBlockHeight() map[byte][]uint64
	RevertShardToBeaconPool(s byte, height uint64)
}

type CrossShardPool interface {
	AddCrossShardBlock(*CrossShardBlock) (map[byte]uint64, byte, error)
	GetValidBlock(map[byte]uint64) map[byte][]*CrossShardBlock
	GetLatestValidBlockHeight() map[byte]uint64
	GetValidBlockHeight() map[byte][]uint64
	GetBlockByHeight(_shardID byte, height uint64) *CrossShardBlock
	RemoveBlockByHeight(map[byte]uint64)
	UpdatePool() map[byte]uint64
	GetAllBlockHeight() map[byte][]uint64
	RevertCrossShardPool(uint64)
	FindBeaconHeightForCrossShardBlock(beaconHeight uint64, fromShardID byte, crossShardBlockHeight uint64) (uint64, error)
}

type ShardPool interface {
	RemoveBlock(height uint64)
	AddShardBlock(block *ShardBlock) error
	GetValidBlockHash() []common.Hash
	GetValidBlock() []*ShardBlock
	GetValidBlockHeight() []uint64
	GetLatestValidBlockHeight() uint64
	SetShardState(height uint64)
	RevertShardPool(uint64)
	GetAllBlockHeight() []uint64
	GetPendingBlockHeight() []uint64
	Start(chan struct{})
}

type BeaconPool interface {
	RemoveBlock(height uint64)
	AddBeaconBlock(block *BeaconBlock) error
	GetValidBlock() []*BeaconBlock
	GetValidBlockHeight() []uint64
	SetBeaconState(height uint64)
	GetBeaconState() uint64
	RevertBeconPool(height uint64)
	GetAllBlockHeight() []uint64
	Start(chan struct{})
	GetPendingBlockHeight() []uint64
}
type TxPool interface {
	// LastUpdated returns the last time a transaction was added to or
	// removed from the source pool.
	LastUpdated() time.Time
	// MiningDescs returns a slice of mining descriptors for all the
	// transactions in the source pool.
	MiningDescs() []*metadata.TxDesc
	// HaveTransaction returns whether or not the passed transaction hash
	// exists in the source pool.
	HaveTransaction(hash *common.Hash) bool
	// RemoveTx remove tx from tx resource
	RemoveTx(txs []metadata.Transaction, isInBlock bool)
	RemoveCandidateList([]string)
	EmptyPool() bool
	MaybeAcceptTransactionForBlockProducing(metadata.Transaction, int64) (*metadata.TxDesc, error)
	MaybeAcceptBatchTransactionForBlockProducing(byte, []metadata.Transaction, int64) ([]*metadata.TxDesc, error)
	//CheckTransactionFee
	// CheckTransactionFee(tx metadata.Transaction) (uint64, error)
	// Check tx validate by it self
	// ValidateTxByItSelf(tx metadata.Transaction) bool
}

type FeeEstimator interface {
	RegisterBlock(block *ShardBlock) error
}

type BestStateInterface interface {
	GetLastBlockTimeStamp() uint64
	GetBlkMinInterval() time.Duration
	GetBlkMaxCreateTime() time.Duration
	CurrentHeight() uint64
	GetCommittee() []string
	GetLastProposerIdx() int
}
