// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	instruction "github.com/incognitochain/incognito-chain/instruction"
	mock "github.com/stretchr/testify/mock"

	signaturecounter "github.com/incognitochain/incognito-chain/blockchain/signaturecounter"
)

// SwapRule is an autogenerated mock type for the SwapRule type
type SwapRule struct {
	mock.Mock
}

// AssignOffset provides a mock function with given fields: lenSubstitute, lenCommittees, numberOfFixedValidators, minCommitteeSize
func (_m *SwapRule) AssignOffset(lenSubstitute int, lenCommittees int, numberOfFixedValidators int, minCommitteeSize int) int {
	ret := _m.Called(lenSubstitute, lenCommittees, numberOfFixedValidators, minCommitteeSize)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int, int, int) int); ok {
		r0 = rf(lenSubstitute, lenCommittees, numberOfFixedValidators, minCommitteeSize)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GenInstructions provides a mock function with given fields: shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty
func (_m *SwapRule) GenInstructions(shardID byte, committees []string, substitutes []string, minCommitteeSize int, maxCommitteeSize int, typeIns int, numberOfFixedValidators int, dcsMaxCommitteeSize int, dcsMinCommitteeSize int, penalty map[string]signaturecounter.Penalty) (*instruction.SwapShardInstruction, []string, []string, []string, []string) {
	ret := _m.Called(shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty)

	var r0 *instruction.SwapShardInstruction
	if rf, ok := ret.Get(0).(func(byte, []string, []string, int, int, int, int, int, int, map[string]signaturecounter.Penalty) *instruction.SwapShardInstruction); ok {
		r0 = rf(shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*instruction.SwapShardInstruction)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func(byte, []string, []string, int, int, int, int, int, int, map[string]signaturecounter.Penalty) []string); ok {
		r1 = rf(shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 []string
	if rf, ok := ret.Get(2).(func(byte, []string, []string, int, int, int, int, int, int, map[string]signaturecounter.Penalty) []string); ok {
		r2 = rf(shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).([]string)
		}
	}

	var r3 []string
	if rf, ok := ret.Get(3).(func(byte, []string, []string, int, int, int, int, int, int, map[string]signaturecounter.Penalty) []string); ok {
		r3 = rf(shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).([]string)
		}
	}

	var r4 []string
	if rf, ok := ret.Get(4).(func(byte, []string, []string, int, int, int, int, int, int, map[string]signaturecounter.Penalty) []string); ok {
		r4 = rf(shardID, committees, substitutes, minCommitteeSize, maxCommitteeSize, typeIns, numberOfFixedValidators, dcsMaxCommitteeSize, dcsMinCommitteeSize, penalty)
	} else {
		if ret.Get(4) != nil {
			r4 = ret.Get(4).([]string)
		}
	}

	return r0, r1, r2, r3, r4
}
