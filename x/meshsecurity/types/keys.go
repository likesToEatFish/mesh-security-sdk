package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name.
	ModuleName = "meshsecurity"

	// StoreKey defines the primary module store key.
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "memory:meshsecurity"

	// RouterKey is the message route
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key.
	QuerierRoute = ModuleName
)

var (
	// ParamsKey is the prefix for the module parameters
	ParamsKey                     = []byte{0x1}
	MaxCapLimitKeyPrefix          = []byte{0x2}
	TotalDelegatedAmountKeyPrefix = []byte{0x3}
	SchedulerKeyPrefix            = []byte{0x4}

	PipedValsetPrefix = []byte{0x5}
)

type PipedValsetOperation byte

const (
	ValsetOperationUndefined PipedValsetOperation = iota
	ValidatorBonded
	ValidatorUnbonded
	ValidatorJailed
	ValidatorTombstoned
	ValidatorUnjailed
	ValidatorModified
	ValidatorSlashed
)

// BuildMaxCapLimitKey build max cap limit store key
func BuildMaxCapLimitKey(contractAddr sdk.AccAddress) []byte {
	return append(MaxCapLimitKeyPrefix, contractAddr.Bytes()...)
}

// BuildTotalDelegatedAmountKey build delegated amount store key for given contract
func BuildTotalDelegatedAmountKey(contractAddr sdk.AccAddress) []byte {
	return append(TotalDelegatedAmountKeyPrefix, contractAddr.Bytes()...)
}

// BuildSchedulerTypeKeyPrefix internal scheduler store key
func BuildSchedulerTypeKeyPrefix(tp SchedulerTaskType) ([]byte, error) {
	if tp == SchedulerTaskUndefined {
		return nil, ErrInvalid.Wrapf("scheduler type: %x", tp)
	}
	return append(SchedulerKeyPrefix, byte(tp)), nil
}

// BuildSchedulerHeightKeyPrefix build store key prefix
func BuildSchedulerHeightKeyPrefix(tp SchedulerTaskType, blockHeight uint64) ([]byte, error) {
	prefix, err := BuildSchedulerTypeKeyPrefix(tp)
	if err != nil {
		return nil, err
	}
	return append(prefix, sdk.Uint64ToBigEndian(blockHeight)...), nil
}

// BuildSchedulerContractKey build store key
func BuildSchedulerContractKey(tp SchedulerTaskType, blockHeight uint64, contractAddr sdk.AccAddress) ([]byte, error) {
	prefix, err := BuildSchedulerHeightKeyPrefix(tp, blockHeight)
	if err != nil {
		return nil, err
	}
	return append(prefix, contractAddr.Bytes()...), nil
}

// BuildPipedValsetOpKey build store key for the temporary valset operation store
func BuildPipedValsetOpKey(op PipedValsetOperation, val sdk.ValAddress) []byte {
	if op == ValsetOperationUndefined {
		panic("empty operation")
	}
	k := append(append(PipedValsetPrefix, val...), byte(op))
	// return
	return k
}
