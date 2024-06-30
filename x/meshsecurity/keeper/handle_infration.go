package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/osmosis-labs/mesh-security-sdk/x/meshsecurity/types"
)

func (k Keeper) HandleInfration(ctx sdk.Context, infration types.SlashInfo) {
	consAddr := infration.GetConsensusAddress()
	validator := k.Staking.ValidatorByConsAddr(ctx, consAddr)
	if validator == nil || validator.IsUnbonded() {
		return
	}

	if !validator.GetOperator().Empty() {
		if _, err := k.Slashing.GetPubkey(ctx, consAddr.Bytes()); err != nil {
			// Ignore meshsecurity that cannot be handled.
			ModuleLogger(ctx).Error(fmt.Sprintf("ignore meshsecurity; expected public key for validator %s not found", consAddr))
			return
		}
	}

	// ignore if the validator is already tombstoned
	if k.Slashing.IsTombstoned(ctx, consAddr) {
		ModuleLogger(ctx).Info(
			"ignored equivocation; validator already tombstoned",
			"validator", consAddr,
			"infraction_height", infration.GetHeight(),
			"infraction_time", infration.GetTime(),
		)
		return
	}

	k.SlashWithInfractionReason(
		ctx,
		infration,
		k.Slashing.SlashFractionDoubleSign(ctx),
		stakingtypes.Infraction_INFRACTION_DOUBLE_SIGN,
	)
	// Jail the validator if not already jailed. This will begin unbonding the
	// validator if not already unbonding (tombstoned).
	if !validator.IsJailed() {
		k.Jail(ctx, consAddr)
	}

	k.Tombstone(ctx, consAddr)
}
