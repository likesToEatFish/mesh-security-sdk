package types

import (
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func FromABCIMeshInfaction(e abci.Misbehavior) SlashInfo {
	bech32PrefixConsAddr := sdk.GetConfig().GetBech32ConsensusAddrPrefix()
	consAddr, err := sdk.Bech32ifyAddressBytes(bech32PrefixConsAddr, e.Validator.Address)
	if err != nil {
		panic(err)
	}

	return SlashInfo{
		InfractionHeight: e.Height,
		Power:            e.Validator.Power,
		TotalSlashAmount: "",
		SlashFraction:    "",
		TimeInfraction:   e.Time,
		ConsensusAddress: consAddr,
	}
}

type SlashInfo struct {
	InfractionHeight int64
	Power            int64
	TotalSlashAmount string
	SlashFraction    string
	TimeInfraction   time.Time
	ConsensusAddress string
}

func (i SlashInfo) GetTime() time.Time {
	return i.TimeInfraction
}

func (i SlashInfo) GetHeight() int64 {
	return i.InfractionHeight
}

func (e SlashInfo) GetConsensusAddress() sdk.ConsAddress {
	addr, _ := sdk.ConsAddressFromBech32(e.ConsensusAddress)
	return addr
}

func (i SlashInfo) GetPower() int64 {
	return i.Power
}
