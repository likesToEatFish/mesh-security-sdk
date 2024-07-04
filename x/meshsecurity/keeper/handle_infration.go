package keeper

import (
	"fmt"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/osmosis-labs/mesh-security-sdk/x/meshsecurity/types"
)

func (k *Keeper) HandleInfration(ctx sdk.Context, infration types.SlashInfo) {
	consAddr := infration.GetConsensusAddress()
	validator := k.Staking.ValidatorByConsAddr(ctx, consAddr)
	if validator == nil || validator.IsUnbonded() {
		return
	}

	// SlashingKeeperDecorator := CaptureTombstoneDecorator(k)
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
		fmt.Println("is tombstoneddddddddddd")
		return
	}

	k.SlashWithInfractionReason(
		ctx,
		infration,
		k.Slashing.SlashFractionDowntime(ctx),
		stakingtypes.Infraction_INFRACTION_DOUBLE_SIGN,
	)
	// Jail the validator if not already jailed. This will begin unbonding the
	// validator if not already unbonding (tombstoned).
	if !validator.IsJailed() {
		k.Jail(ctx, consAddr)
	}

	k.Tombstone(ctx, consAddr)
}

func (k Keeper) HandleEvenSlash(ctx sdk.Context) {
	k.iterateDelegateVirtual(ctx, func(actor sdk.AccAddress, val sdk.ValAddress) bool {
		validator, found := k.Staking.GetValidator(ctx, val)
		if !found {
			return false
		}
		// jail
		// unjail
		if validator.IsJailed() {
			k.Jail1(ctx, val)
		} else {
			k.Unjail1(ctx, val)
		}
		// tombstone

		// slash

		return false
	})
}

func (k Keeper) HandleValidatorSignature(ctx sdk.Context, params slashingtypes.Params, addr cryptotypes.Address, power int64, signed bool) {
	// fetch signing info
	consAddr := sdk.ConsAddress(addr)
	signInfo, found := k.Slashing.GetValidatorSigningInfo(ctx, consAddr)
	if !found {
		panic(fmt.Sprintf("Expected signing info for validator %s but not found", consAddr))
	}
	// params.MinSignedPerWindowInt()
	minSignedPerWindow := params.MinSignedPerWindow.MulInt64(params.SignedBlocksWindow).RoundInt64()

	height := ctx.BlockHeight()

	minHeight := signInfo.StartHeight + params.SignedBlocksWindow
	maxMissed := params.SignedBlocksWindow - minSignedPerWindow
	if height > minHeight && signInfo.MissedBlocksCounter > maxMissed {
		distributionHeight := height - sdk.ValidatorUpdateDelay - 1
		slashInfo := types.SlashInfo{
			Power:            power,
			ConsensusAddress: consAddr.String(),
			InfractionHeight: distributionHeight,
			TimeInfraction:   ctx.WithBlockHeight(distributionHeight).BlockTime(),
		}

		k.SlashWithInfractionReason(ctx, slashInfo, k.Slashing.SlashFractionDowntime(ctx), stakingtypes.Infraction_INFRACTION_DOWNTIME)
		k.Jail(ctx, consAddr)
	}

}
