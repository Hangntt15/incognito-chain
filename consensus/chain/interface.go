package chain

import (
	"time"

	"github.com/incognitochain/incognito-chain/common"
	"github.com/incognitochain/incognito-chain/wire"
)

type ConsensusInterface interface {
	GetInfo() string

	Start()
	Stop()
	IsRun() bool

	ReceiveProposeMsg(interface{})
	ReceivePrepareMsg(interface{})

	ProcessBFTMsg(interface{})
	ValidateBlock(interface{}) error
}

type BlockInterface interface {
	GetHeight() uint64
	Hash() *common.Hash
	AddValidationField(validateData string) error
	GetValidationField() string
	GetRound() int
	GetRoundKey() string
}

type ChainInterface interface {
	GetConsensusEngine() ConsensusInterface
	PushMessageToValidator(wire.Message) error
	GetLastBlockTimeStamp() uint64
	GetBlkMinTime() time.Duration
	IsReady() bool
	GetHeight() uint64
	GetCommitteeSize() int
	GetNodePubKeyCommitteeIndex() int
	GetLastProposerIndex() int
	GetNodePubKey() string
	CreateNewBlock(round int) BlockInterface
	InsertBlk(interface{}, bool)
	ValidateBlock(interface{}) error
	ValidatePreSignBlock(interface{}) error
	GetActiveShardNumber() int
}

type Node interface {
	PushMessageToShard(wire.Message, byte) error
	PushMessageToBeacon(wire.Message) error
	IsEnableMining() bool
}
