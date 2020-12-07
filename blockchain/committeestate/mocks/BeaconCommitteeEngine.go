// Code generated by mockery v2.0.0. DO NOT EDIT.

package mocks

import (
	committeestate "github.com/incognitochain/incognito-chain/blockchain/committeestate"
	common "github.com/incognitochain/incognito-chain/common"

	incognitokey "github.com/incognitochain/incognito-chain/incognitokey"

	instruction "github.com/incognitochain/incognito-chain/instruction"

	mock "github.com/stretchr/testify/mock"

	privacy "github.com/incognitochain/incognito-chain/privacy"
)

// BeaconCommitteeEngine is an autogenerated mock type for the BeaconCommitteeEngine type
type BeaconCommitteeEngine struct {
	mock.Mock
}

// AbortUncommittedBeaconState provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) AbortUncommittedBeaconState() {
	_m.Called()
}

// ActiveShards provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) ActiveShards() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Clone provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) Clone() committeestate.BeaconCommitteeEngine {
	ret := _m.Called()

	var r0 committeestate.BeaconCommitteeEngine
	if rf, ok := ret.Get(0).(func() committeestate.BeaconCommitteeEngine); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(committeestate.BeaconCommitteeEngine)
		}
	}

	return r0
}

// Commit provides a mock function with given fields: _a0
func (_m *BeaconCommitteeEngine) Commit(_a0 *committeestate.BeaconCommitteeStateHash) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*committeestate.BeaconCommitteeStateHash) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateAllSwapShardInstructions provides a mock function with given fields: env
func (_m *BeaconCommitteeEngine) GenerateAllSwapShardInstructions(env *committeestate.BeaconCommitteeStateEnvironment) ([]*instruction.SwapShardInstruction, error) {
	ret := _m.Called(env)

	var r0 []*instruction.SwapShardInstruction
	if rf, ok := ret.Get(0).(func(*committeestate.BeaconCommitteeStateEnvironment) []*instruction.SwapShardInstruction); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*instruction.SwapShardInstruction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*committeestate.BeaconCommitteeStateEnvironment) error); ok {
		r1 = rf(env)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GenerateAssignInstruction provides a mock function with given fields: rand, assignOffset, activeShards
func (_m *BeaconCommitteeEngine) GenerateAssignInstruction(rand int64, assignOffset int, activeShards int) ([]*instruction.AssignInstruction, []string, map[byte][]string) {
	ret := _m.Called(rand, assignOffset, activeShards)

	var r0 []*instruction.AssignInstruction
	if rf, ok := ret.Get(0).(func(int64, int, int) []*instruction.AssignInstruction); ok {
		r0 = rf(rand, assignOffset, activeShards)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*instruction.AssignInstruction)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(int64, int, int) []string); ok {
		r1 = rf(rand, assignOffset, activeShards)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 map[byte][]string
	if rf, ok := ret.Get(2).(func(int64, int, int) map[byte][]string); ok {
		r2 = rf(rand, assignOffset, activeShards)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(map[byte][]string)
		}
	}

	return r0, r1, r2
}

// GetAllCandidateSubstituteCommittee provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetAllCandidateSubstituteCommittee() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// GetAutoStaking provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetAutoStaking() map[string]bool {
	ret := _m.Called()

	var r0 map[string]bool
	if rf, ok := ret.Get(0).(func() map[string]bool); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]bool)
		}
	}

	return r0
}

// GetBeaconCommittee provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetBeaconCommittee() []incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() []incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetBeaconHash provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetBeaconHash() common.Hash {
	ret := _m.Called()

	var r0 common.Hash
	if rf, ok := ret.Get(0).(func() common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(common.Hash)
		}
	}

	return r0
}

// GetBeaconHeight provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetBeaconHeight() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// GetBeaconSubstitute provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetBeaconSubstitute() []incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() []incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetCandidateBeaconWaitingForCurrentRandom provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetCandidateBeaconWaitingForCurrentRandom() []incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() []incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetCandidateBeaconWaitingForNextRandom provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetCandidateBeaconWaitingForNextRandom() []incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() []incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetCandidateShardWaitingForCurrentRandom provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetCandidateShardWaitingForCurrentRandom() []incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() []incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetCandidateShardWaitingForNextRandom provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetCandidateShardWaitingForNextRandom() []incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() []incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetOneShardCommittee provides a mock function with given fields: shardID
func (_m *BeaconCommitteeEngine) GetOneShardCommittee(shardID byte) []incognitokey.CommitteePublicKey {
	ret := _m.Called(shardID)

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func(byte) []incognitokey.CommitteePublicKey); ok {
		r0 = rf(shardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetOneShardSubstitute provides a mock function with given fields: shardID
func (_m *BeaconCommitteeEngine) GetOneShardSubstitute(shardID byte) []incognitokey.CommitteePublicKey {
	ret := _m.Called(shardID)

	var r0 []incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func(byte) []incognitokey.CommitteePublicKey); ok {
		r0 = rf(shardID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetRewardReceiver provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetRewardReceiver() map[string]privacy.PaymentAddress {
	ret := _m.Called()

	var r0 map[string]privacy.PaymentAddress
	if rf, ok := ret.Get(0).(func() map[string]privacy.PaymentAddress); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]privacy.PaymentAddress)
		}
	}

	return r0
}

// GetShardCommittee provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetShardCommittee() map[byte][]incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 map[byte][]incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() map[byte][]incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[byte][]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetShardSubstitute provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetShardSubstitute() map[byte][]incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 map[byte][]incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() map[byte][]incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[byte][]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// GetStakingTx provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetStakingTx() map[string]common.Hash {
	ret := _m.Called()

	var r0 map[string]common.Hash
	if rf, ok := ret.Get(0).(func() map[string]common.Hash); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]common.Hash)
		}
	}

	return r0
}

// GetUncommittedCommittee provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) GetUncommittedCommittee() map[byte][]incognitokey.CommitteePublicKey {
	ret := _m.Called()

	var r0 map[byte][]incognitokey.CommitteePublicKey
	if rf, ok := ret.Get(0).(func() map[byte][]incognitokey.CommitteePublicKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[byte][]incognitokey.CommitteePublicKey)
		}
	}

	return r0
}

// InitCommitteeState provides a mock function with given fields: env
func (_m *BeaconCommitteeEngine) InitCommitteeState(env *committeestate.BeaconCommitteeStateEnvironment) {
	_m.Called(env)
}

// SplitReward provides a mock function with given fields: _a0
func (_m *BeaconCommitteeEngine) SplitReward(_a0 *committeestate.BeaconCommitteeStateEnvironment) (map[common.Hash]uint64, map[common.Hash]uint64, map[common.Hash]uint64, map[common.Hash]uint64, error) {
	ret := _m.Called(_a0)

	var r0 map[common.Hash]uint64
	if rf, ok := ret.Get(0).(func(*committeestate.BeaconCommitteeStateEnvironment) map[common.Hash]uint64); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[common.Hash]uint64)
		}
	}

	var r1 map[common.Hash]uint64
	if rf, ok := ret.Get(1).(func(*committeestate.BeaconCommitteeStateEnvironment) map[common.Hash]uint64); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[common.Hash]uint64)
		}
	}

	var r2 map[common.Hash]uint64
	if rf, ok := ret.Get(2).(func(*committeestate.BeaconCommitteeStateEnvironment) map[common.Hash]uint64); ok {
		r2 = rf(_a0)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(map[common.Hash]uint64)
		}
	}

	var r3 map[common.Hash]uint64
	if rf, ok := ret.Get(3).(func(*committeestate.BeaconCommitteeStateEnvironment) map[common.Hash]uint64); ok {
		r3 = rf(_a0)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(map[common.Hash]uint64)
		}
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(*committeestate.BeaconCommitteeStateEnvironment) error); ok {
		r4 = rf(_a0)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// UpdateCommitteeState provides a mock function with given fields: env
func (_m *BeaconCommitteeEngine) UpdateCommitteeState(env *committeestate.BeaconCommitteeStateEnvironment) (*committeestate.BeaconCommitteeStateHash, *committeestate.CommitteeChange, [][]string, error) {
	ret := _m.Called(env)

	var r0 *committeestate.BeaconCommitteeStateHash
	if rf, ok := ret.Get(0).(func(*committeestate.BeaconCommitteeStateEnvironment) *committeestate.BeaconCommitteeStateHash); ok {
		r0 = rf(env)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*committeestate.BeaconCommitteeStateHash)
		}
	}

	var r1 *committeestate.CommitteeChange
	if rf, ok := ret.Get(1).(func(*committeestate.BeaconCommitteeStateEnvironment) *committeestate.CommitteeChange); ok {
		r1 = rf(env)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*committeestate.CommitteeChange)
		}
	}

	var r2 [][]string
	if rf, ok := ret.Get(2).(func(*committeestate.BeaconCommitteeStateEnvironment) [][]string); ok {
		r2 = rf(env)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([][]string)
		}
	}

	var r3 error
	if rf, ok := ret.Get(3).(func(*committeestate.BeaconCommitteeStateEnvironment) error); ok {
		r3 = rf(env)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// Version provides a mock function with given fields:
func (_m *BeaconCommitteeEngine) Version() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
