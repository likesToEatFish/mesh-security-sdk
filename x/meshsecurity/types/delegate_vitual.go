package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDelegationVirtual(deAdd string, actor string, valAdd string, amount math.Int) DelegationVirtual {
	return DelegationVirtual{
		DelegatorAddress: deAdd,
		ValidatorAddress: valAdd,
		Amount:           amount,
	}
}

func (d DelegationVirtual) GetValidator() (sdk.ValAddress, error) {
	valAddr, err := sdk.ValAddressFromBech32(d.ValidatorAddress)
	if err != nil {
		return sdk.ValAddress{}, err
	}
	return valAddr, nil
}

func (d DelegationVirtual) GetDelegation() (sdk.AccAddress, error) {
	delAddr, err := sdk.AccAddressFromBech32(d.DelegatorAddress)
	if err != nil {
		return sdk.AccAddress{}, err
	}
	return delAddr, nil
}
