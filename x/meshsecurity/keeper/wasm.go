package keeper

import (
	"encoding/json"
	// "fmt"

	errorsmod "cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/mesh-security-sdk/x/meshsecurity/contract"
)

// SendHandleEpoch send epoch handling message to virtual staking contract via sudo
func (k Keeper) SendHandleEpoch(ctx sdk.Context, contractAddr sdk.AccAddress) error {
	msg := contract.SudoMsg{
		HandleEpoch: &struct{}{},
	}
	return k.doSudoCall(ctx, contractAddr, msg)
}

// SendValsetUpdate submit the valset update report to the virtual staking contract via sudo
func (k Keeper) SendValsetUpdate(ctx sdk.Context, contractAddr sdk.AccAddress, v contract.ValsetUpdate) error {
	msg := contract.SudoMsg{
		ValsetUpdate: &v,
	}
	return k.doSudoCall(ctx, contractAddr, msg)
}

// caller must ensure gas limits are set proper and handle panics
func (k Keeper) doSudoCall(ctx sdk.Context, contractAddr sdk.AccAddress, msg contract.SudoMsg) error {
	bz, err := json.Marshal(msg)
	// fmt.Println("msg:", msg)
	if err != nil {
		return errorsmod.Wrap(err, "marshal sudo msg")
	}
	// {30E2301C8C6F801FC1D4218AF4AE03509B745725169F7A379C820C167957E363 Error parsing into type mesh_apis::virtual_staking_api::SudoMsg: Invalid number.: execute wasm contract failed [!cosm!wasm/wasmd@v0.45.0/x/wasm/keeper/keeper.go:518] <nil> <nil> 69236 500000 0}
	_, err = k.wasm.Sudo(ctx, contractAddr, bz)
	return err
}
