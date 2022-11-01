package cli_test

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"point/testutil/network"
	"point/testutil/nullify"
	"point/x/lockenomics/client/cli"
	"point/x/lockenomics/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithDelegationLockObjects(t *testing.T, n int) (*network.Network, []types.DelegationLock) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	delegators := [7]string{"cosmos156gqf9837u7d4c4678yt3rl4ls9c5vuuxyhkw6",
		"cosmos14lultfckehtszvzw4ehu0apvsr77afvyhgqhwh",
		"cosmos19lss6zgdh5vvcpjhfftdghrpsw7a4434ut4md0",
		"cosmos196ax4vc0lwpxndu9dyhvca7jhxp70rmcfhxsrt",
		"cosmos1z8zjv3lntpwxua0rtpvgrcwl0nm0tltgyuy0nd",
		"cosmos1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqda85ee",
		"cosmos1tflk30mq5vgqjdly92kkhhq3raev2hnzldd74z",
	}
	validators := [7]string{"cosmosvaloper156gqf9837u7d4c4678yt3rl4ls9c5vuursrrzf",
		"cosmosvaloper14lultfckehtszvzw4ehu0apvsr77afvyju5zzy",
		"cosmosvaloper19lss6zgdh5vvcpjhfftdghrpsw7a4434elpwpu",
		"cosmosvaloper196ax4vc0lwpxndu9dyhvca7jhxp70rmcvrj90c",
		"cosmosvaloper1z8zjv3lntpwxua0rtpvgrcwl0nm0tltgpgs6l7",
		"cosmosvaloper1qaa9zej9a0ge3ugpx3pxyx602lxh3ztqgfnp42",
		"cosmosvaloper1tflk30mq5vgqjdly92kkhhq3raev2hnz6eete3",
	}

	indexes := [len(delegators)]string{}
	for i := 0; i < len(delegators); i++ {
		delegator, _ := sdk.AccAddressFromBech32(delegators[i])
		validator, _ := sdk.ValAddressFromBech32(validators[i])
		key := types.GetDelegationLockKeyString(delegator, validator)
		indexes[i] = key
	}

	for i := 0; i < n; i++ {
		delegationLock := types.DelegationLock{
			Index:     indexes[i],
			Start:     uint64(time.Now().Unix()),
			Length:    12978,
			Delegator: delegators[i],
			Validator: validators[i],
		}
		nullify.Fill(&delegationLock)
		state.DelegationLockList = append(state.DelegationLockList, delegationLock)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.DelegationLockList
}

func TestShowDelegationLock(t *testing.T) {
	net, objs := networkWithDelegationLockObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.DelegationLock
	}{
		{
			desc:    "found",
			idIndex: objs[0].Index,

			args: common,
			obj:  objs[0],
		},
		{
			desc:    "not found",
			idIndex: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idIndex,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowDelegationLock(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetDelegationLockResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.DelegationLock)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.DelegationLock),
				)
			}
		})
	}
}

func TestListDelegationLock(t *testing.T) {
	net, objs := networkWithDelegationLockObjects(t, 5)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDelegationLock(), args)
			require.NoError(t, err)
			var resp types.QueryAllDelegationLockResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DelegationLock), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DelegationLock),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDelegationLock(), args)
			require.NoError(t, err)
			var resp types.QueryAllDelegationLockResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DelegationLock), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.DelegationLock),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDelegationLock(), args)
		require.NoError(t, err)
		var resp types.QueryAllDelegationLockResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.DelegationLock),
		)
	})
}
