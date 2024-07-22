package meshsecurity

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/mesh-security-sdk/x/meshsecurity/keeper"
	"github.com/osmosis-labs/mesh-security-sdk/x/meshsecurity/types"

	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/cosmos/cosmos-sdk/x/slashing/testutil"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtestutil "github.com/cosmos/cosmos-sdk/x/staking/testutil"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestEndBlocker(t *testing.T) {
	var (
		capturedCalls []capturedSudo
		contractErr   error
		logRecords    bytes.Buffer
	)
	pCtx, keepers := keeper.CreateDefaultTestInput(t, keeper.WithWasmKeeperDecorated(func(original types.WasmKeeper) types.WasmKeeper {
		return captureSudos(&capturedCalls, &contractErr)
	}))
	val1 := keeper.MinValidatorFixture(t)
	keepers.StakingKeeper.SetValidator(pCtx, val1)
	k := keepers.MeshKeeper
	var (
		myError             = errors.New("my test error")
		myContractAddr      = sdk.AccAddress(bytes.Repeat([]byte{1}, 32))
		myOtherContractAddr = sdk.AccAddress(bytes.Repeat([]byte{2}, 32))
	)

	specs := map[string]struct {
		setup  func(t *testing.T, ctx sdk.Context)
		assert func(t *testing.T, ctx sdk.Context)
	}{
		"rebalance - multiple contracts": {
			setup: func(t *testing.T, ctx sdk.Context) {
				require.NoError(t,
					k.ScheduleRepeatingTask(ctx, types.SchedulerTaskHandleEpoch, myContractAddr, uint64(ctx.BlockHeight())))
				require.NoError(t,
					k.ScheduleRepeatingTask(ctx, types.SchedulerTaskHandleEpoch, myOtherContractAddr, uint64(ctx.BlockHeight())))
			},
			assert: func(t *testing.T, ctx sdk.Context) {
				require.Len(t, capturedCalls, 2)
				assert.Equal(t, myContractAddr, capturedCalls[0].contractAddress)
				assert.JSONEq(t, `{"handle_epoch":{}}`, string(capturedCalls[0].msg))
				assert.Equal(t, myOtherContractAddr, capturedCalls[1].contractAddress)
				assert.JSONEq(t, `{"handle_epoch":{}}`, string(capturedCalls[1].msg))
				assert.NotContains(t, logRecords.String(), "failed")
			},
		},
		"rebalance - contract errored": {
			setup: func(t *testing.T, ctx sdk.Context) {
				contractErr = myError
				require.NoError(t,
					k.ScheduleRepeatingTask(ctx, types.SchedulerTaskHandleEpoch, myContractAddr, uint64(ctx.BlockHeight())))
				require.NoError(t,
					k.ScheduleRepeatingTask(ctx, types.SchedulerTaskHandleEpoch, myOtherContractAddr, uint64(ctx.BlockHeight())))
			},
			assert: func(t *testing.T, ctx sdk.Context) {
				require.Len(t, capturedCalls, 2)
				assert.Contains(t, logRecords.String(), "failed to execute scheduled task")
			},
		},
		"valset update - multiple contracts": {
			setup: func(t *testing.T, ctx sdk.Context) {
				anyLimit := sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1_000_000_000))
				require.NoError(t, k.SetMaxCapLimit(ctx, myContractAddr, anyLimit))
				require.NoError(t, k.SetMaxCapLimit(ctx, myOtherContractAddr, anyLimit))
				require.NoError(t, k.Hooks().AfterValidatorBonded(ctx, nil, val1.GetOperator()))
			},
			assert: func(t *testing.T, ctx sdk.Context) {
				require.Len(t, capturedCalls, 2)
				assert.Equal(t, myContractAddr, capturedCalls[0].contractAddress)
				exp := fmt.Sprintf(`{"valset_update":{"additions":[{"address":"%s","commission":"0.000000000000000000","max_commission":"0.000000000000000000","max_change_rate":"0.000000000000000000"}],"removals":[],"updated":[],"jailed":[],"unjailed":[],"slashed":[],"tombstoned":[]}}`, val1.GetOperator())
				assert.JSONEq(t, exp, string(capturedCalls[0].msg))

				assert.Equal(t, myOtherContractAddr, capturedCalls[1].contractAddress)
				assert.JSONEq(t, exp, string(capturedCalls[1].msg))
				assert.NotContains(t, logRecords.String(), "failed")
			},
		},
		"valset update - contract errored": {
			setup: func(t *testing.T, ctx sdk.Context) {
				anyLimit := sdk.NewCoin(sdk.DefaultBondDenom, sdkmath.NewInt(1_000_000_000))
				require.NoError(t, k.SetMaxCapLimit(ctx, myContractAddr, anyLimit))
				require.NoError(t, k.SetMaxCapLimit(ctx, myOtherContractAddr, anyLimit))
				require.NoError(t, k.Hooks().AfterValidatorBonded(ctx, nil, val1.GetOperator()))
				contractErr = myError
			},
			assert: func(t *testing.T, ctx sdk.Context) {
				require.Len(t, capturedCalls, 2)
				assert.Contains(t, logRecords.String(), "failed to execute scheduled task")
			},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			capturedCalls, contractErr = nil, nil
			logRecords.Reset()
			ctx, _ := pCtx.CacheContext()
			spec.setup(t, ctx)
			// when
			EndBlocker(ctx.WithLogger(log.NewTMLogger(log.NewSyncWriter(&logRecords))), k, DefaultExecutionResponseHandler())
			spec.assert(t, ctx)
		})
	}
}

type capturedSudo = struct {
	contractAddress sdk.AccAddress
	msg             []byte
}

func captureSudos(captured *[]capturedSudo, e *error) *keeper.MockWasmKeeper {
	return &keeper.MockWasmKeeper{
		SudoFn: func(ctx sdk.Context, contractAddress sdk.AccAddress, msg []byte) ([]byte, error) {
			*captured = append(*captured, capturedSudo{contractAddress: contractAddress, msg: msg})
			return nil, *e
		},
		HasContractInfoFn: func(ctx sdk.Context, contractAddress sdk.AccAddress) bool {
			return true
		},
	}
}

// Trình xác thực nguồn mặc định được khởi tạo để có trong các thử nghiệm
// The default power validators are initialized to have within tests
var InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)

func TestBeginBlocker(t *testing.T) {
	var interfaceRegistry codectypes.InterfaceRegistry
	var bankKeeper bankkeeper.Keeper
	var stakingKeeper *stakingkeeper.Keeper
	var slashingKeeper slashingkeeper.Keeper

	app, err := simtestutil.Setup(
		testutil.AppConfig,
		&interfaceRegistry,
		&bankKeeper,
		&stakingKeeper,
		&slashingKeeper,
	)
	require.NoError(t, err)

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	pks := simtestutil.CreateTestPubKeys(1)
	simtestutil.AddTestAddrsFromPubKeys(bankKeeper, stakingKeeper, ctx, pks, stakingKeeper.TokensFromConsensusPower(ctx, 200))
	addr, pk := sdk.ValAddress(pks[0].Address()), pks[0]
	tstaking := stakingtestutil.NewHelper(t, ctx, stakingKeeper)

	// bond the validator
	power := int64(100)
	amt := tstaking.CreateValidatorWithValPower(addr, pk, power, true)
	staking.EndBlocker(ctx, stakingKeeper)
	require.Equal(
		t, bankKeeper.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.NewCoins(sdk.NewCoin(stakingKeeper.GetParams(ctx).BondDenom, InitTokens.Sub(amt))),
	)
	require.Equal(t, amt, stakingKeeper.Validator(ctx, addr).GetBondedTokens())

	val := abci.Validator{
		Address: pk.Address(),
		Power:   power,
	}

	// mark the validator as having signed
	req := abci.RequestBeginBlock{
		LastCommitInfo: abci.CommitInfo{
			Votes: []abci.VoteInfo{{
				Validator:       val,
				SignedLastBlock: true,
			}},
		},
	}

	slashing.BeginBlocker(ctx, req, slashingKeeper)

	info, found := slashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(pk.Address()))
	require.True(t, found)
	require.Equal(t, ctx.BlockHeight(), info.StartHeight)
	require.Equal(t, int64(1), info.IndexOffset)
	require.Equal(t, time.Unix(0, 0).UTC(), info.JailedUntil)
	require.Equal(t, int64(0), info.MissedBlocksCounter)

	height := int64(0)

	// for 1000 blocks, mark the validator as having signed
	for ; height < slashingKeeper.SignedBlocksWindow(ctx); height++ {
		ctx = ctx.WithBlockHeight(height)
		req = abci.RequestBeginBlock{
			LastCommitInfo: abci.CommitInfo{
				Votes: []abci.VoteInfo{{
					Validator:       val,
					SignedLastBlock: true,
				}},
			},
		}

		slashing.BeginBlocker(ctx, req, slashingKeeper)
	}

	// for 500 blocks, mark the validator as having not signed
	for ; height < ((slashingKeeper.SignedBlocksWindow(ctx) * 2) - slashingKeeper.MinSignedPerWindow(ctx) + 1); height++ {
		ctx = ctx.WithBlockHeight(height)
		req = abci.RequestBeginBlock{
			LastCommitInfo: abci.CommitInfo{
				Votes: []abci.VoteInfo{{
					Validator:       val,
					SignedLastBlock: false,
				}},
			},
		}

		slashing.BeginBlocker(ctx, req, slashingKeeper)
	}

	// end block
	staking.EndBlocker(ctx, stakingKeeper)

	// validator should be jailed
	validator, found := stakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(pk))
	require.True(t, found)
	require.Equal(t, stakingtypes.Unbonding, validator.GetStatus())
}
