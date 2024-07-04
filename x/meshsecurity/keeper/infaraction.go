package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/osmosis-labs/mesh-security-sdk/x/meshsecurity/types"
)

// SlashWithInfractionReason implementation doesn't require the infraction (types.Infraction) to work but is required by Interchain Security.
func (k Keeper) SlashWithInfractionReason(ctx sdk.Context, infraction types.SlashInfo, slashFactor sdk.Dec, _ stakingtypes.Infraction) {
	consAddr := infraction.GetConsensusAddress()
	power := infraction.GetPower()
	infractionHeight := infraction.GetHeight()
	val := k.Staking.ValidatorByConsAddr(ctx, consAddr)

	totalSlashAmount := k.Staking.SlashWithInfractionReason(ctx, consAddr, power, infractionHeight, slashFactor, stakingtypes.Infraction_INFRACTION_DOUBLE_SIGN)
	if val == nil {
		ModuleLogger(ctx).
			Error("can not propagate slash: validator not found", "validator", consAddr.String())
	} else if err := k.ScheduleSlashed(ctx, infraction, val.GetOperator(), totalSlashAmount, slashFactor); err != nil {
		ModuleLogger(ctx).
			Error("can not propagate slash: schedule event",
				"cause", err,
				"validator", consAddr.String())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSlash,
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)
}

// Jail captures the jail event and calls the decorated staking keeper jail method
func (k Keeper) Jail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	val := k.Staking.ValidatorByConsAddr(ctx, consAddr)
	if val == nil {
		ModuleLogger(ctx).
			Error("can not propagate jail: validator not found", "validator", consAddr.String())
	} else if err := k.ScheduleJailed(ctx, val.GetOperator()); err != nil {
		ModuleLogger(ctx).
			Error("can not propagate jail: schedule event",
				"cause", err,
				"validator", consAddr.String())
	}
	// k.Staking.Jail(ctx, consAddr)
}

// Unjail captures the unjail event and calls the decorated staking keeper unjail method
func (k Keeper) Unjail(ctx sdk.Context, consAddr sdk.ConsAddress) {
	val := k.Staking.ValidatorByConsAddr(ctx, consAddr)
	if val == nil {
		ModuleLogger(ctx).
			Error("can not propagate unjail: validator not found", "validator", consAddr.String())
	} else if err := k.ScheduleUnjailed(ctx, val.GetOperator()); err != nil {
		ModuleLogger(ctx).
			Error("can not propagate unjail: schedule event",
				"cause", err,
				"validator", consAddr.String())
	}
	// k.Staking.Unjail(ctx, consAddr)
}

// Tombstone is executed in the end-blocker by the evidence module
func (k Keeper) Tombstone(ctx sdk.Context, address sdk.ConsAddress) {
	v, ok := k.Staking.GetValidatorByConsAddr(ctx, address)
	if !ok {
		ModuleLogger(ctx).
			Error("can not propagate tompstone: validator not found", "validator", address.String())
	} else if err := k.ScheduleTombstoned(ctx, v.GetOperator()); err != nil {
		ModuleLogger(ctx).
			Error("can not propagate tompstone: scheduler",
				"cause", err,
				"validator", address.String())
	}
	// k.Slashing.Tombstone(ctx, address)
}
