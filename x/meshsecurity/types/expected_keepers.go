package types

import (
	"time"

	"cosmossdk.io/math"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type SDKBankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amounts sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type XBankKeeper interface {
	SDKBankKeeper
	AddSupplyOffset(ctx sdk.Context, denom string, offsetAmount math.Int)
}

// SDKStakingKeeper expected staking keeper.
type SDKStakingKeeper interface {
	BondDenom(ctx sdk.Context) string
	GetAllValidators(ctx sdk.Context) (validators []stakingtypes.Validator)
	GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator stakingtypes.Validator, found bool)
	ValidateUnbondAmount(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, amt math.Int) (shares sdk.Dec, err error)
	Delegate(ctx sdk.Context, delAddr sdk.AccAddress, bondAmt math.Int, tokenSrc stakingtypes.BondStatus, validator stakingtypes.Validator, subtractAccount bool) (newShares sdk.Dec, err error)
	GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation stakingtypes.Delegation, found bool)
	UnbondingTime(ctx sdk.Context) time.Duration
	GetParams(ctx sdk.Context) stakingtypes.Params
	Unbond(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, shares sdk.Dec) (amount math.Int, err error)
	IterateBondedValidatorsByPower(ctx sdk.Context, fn func(int64, stakingtypes.ValidatorI) bool)
	TotalBondedTokens(ctx sdk.Context) math.Int
	IterateDelegations(ctx sdk.Context, delegator sdk.AccAddress, fn func(int64, stakingtypes.DelegationI) bool)
	GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (stakingtypes.Validator, bool)
	ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI
	// slash the validator and delegators of the validator, specifying offence height, offence power, and slash fraction
	Slash(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec) math.Int
	SlashWithInfractionReason(sdk.Context, sdk.ConsAddress, int64, int64, sdk.Dec, stakingtypes.Infraction) math.Int
	Jail(sdk.Context, sdk.ConsAddress)   // jail a validator
	Unjail(sdk.Context, sdk.ConsAddress) // unjail a validator
}

type XStakingKeeper interface {
	SDKStakingKeeper
	InstantUndelegate(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress, sharesAmount sdk.Dec) (sdk.Coins, error)
}

// CommunityPoolKeeper expected distribution keeper.
type CommunityPoolKeeper interface {
	WithdrawDelegationRewards(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (sdk.Coins, error)
}

// AccountKeeper interface contains functions for getting accounts and the module address
type AccountKeeper interface {
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}

// WasmKeeper abstract wasm keeper
type WasmKeeper interface {
	Sudo(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error)
	HasContractInfo(ctx sdk.Context, contractAddress sdk.AccAddress) bool
}

type SlashingKeeper interface {
	GetPubkey(sdk.Context, cryptotypes.Address) (cryptotypes.PubKey, error)
	IsTombstoned(sdk.Context, sdk.ConsAddress) bool
	HasValidatorSigningInfo(sdk.Context, sdk.ConsAddress) bool
	Tombstone(sdk.Context, sdk.ConsAddress)
	Slash(sdk.Context, sdk.ConsAddress, sdk.Dec, int64, int64)
	SlashWithInfractionReason(sdk.Context, sdk.ConsAddress, sdk.Dec, int64, int64, stakingtypes.Infraction)
	SlashFractionDoubleSign(sdk.Context) sdk.Dec
	SlashFractionDowntime(sdk.Context) sdk.Dec
	Jail(sdk.Context, sdk.ConsAddress)
	JailUntil(sdk.Context, sdk.ConsAddress, time.Time)
	GetParams(sdk.Context) slashingtypes.Params
	GetValidatorSigningInfo(sdk.Context, sdk.ConsAddress) (info slashingtypes.ValidatorSigningInfo, found bool)
}
