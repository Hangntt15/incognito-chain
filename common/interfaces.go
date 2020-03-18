package common

import (
	"math"
)

type BlockPoolInterface interface {
	GetPrevHash() Hash
	Hash() *Hash
	GetHeight() uint64
	GetShardID() int
}

type CrossShardBlkPoolInterface interface {
	Hash() *Hash
	GetHeight() uint64
}

type BlockInterface interface {
	GetVersion() int
	GetHeight() uint64
	Hash() *Hash
	// AddValidationField(validateData string) error
	GetProducer() string
	GetValidationField() string
	GetRound() int
	GetRoundKey() string
	GetInstructions() [][]string
	GetConsensusType() string
	GetCurrentEpoch() uint64
	GetProduceTime() int64
	GetProposeTime() int64
	GetPrevHash() Hash
	GetProposer() string
}

type ChainInterface interface {
	GetShardID() int
}

const TIMESLOT = 2

func CalculateTimeSlot(time int64) int64 {
	return int64(math.Floor(float64(time / TIMESLOT)))
}
