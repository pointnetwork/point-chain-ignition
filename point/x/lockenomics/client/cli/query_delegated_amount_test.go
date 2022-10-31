package cli_test

import (
	"fmt"
	"strconv"
	"testing"

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

func networkWithDelegatedAmountObjects(t *testing.T, n int) (*network.Network, []types.DelegatedAmount) {
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

	for i := 0; i < n; i++ {
		key, _ := types.GetDelegatedAmountKeyStringFromString(delegators[i], validators[i])
		delegatedAmount := types.DelegatedAmount{
			Index:     key,
			Delegator: delegators[i],
			Validator: validators[i],
			Amount:    12387,
			Creator:   delegators[i],
		}
		nullify.Fill(&delegatedAmount)
		state.DelegatedAmountList = append(state.DelegatedAmountList, delegatedAmount)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.DelegatedAmountList
}

func TestShowDelegatedAmount(t *testing.T) {
	net, objs := networkWithDelegatedAmountObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc    string
		idIndex string

		args []string
		err  error
		obj  types.DelegatedAmount
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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowDelegatedAmount(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetDelegatedAmountResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.DelegatedAmount)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.DelegatedAmount),
				)
			}
		})
	}
}

func TestListDelegatedAmount(t *testing.T) {
	net, objs := networkWithDelegatedAmountObjects(t, 5)

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
		step := 7
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDelegatedAmount(), args)
			require.NoError(t, err)
			var resp types.QueryAllDelegatedAmountResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DelegatedAmount), step)
			require.Subset(t,
				nullify.Fill(resp.DelegatedAmount),
				nullify.Fill(objs),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 7
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDelegatedAmount(), args)
			require.NoError(t, err)
			var resp types.QueryAllDelegatedAmountResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.DelegatedAmount), step)
			require.Subset(t,
				nullify.Fill(resp.DelegatedAmount),
				nullify.Fill(objs),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)+1), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListDelegatedAmount(), args)
		require.NoError(t, err)
		var resp types.QueryAllDelegatedAmountResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		//Adding 1 to number of objects because 1 additional entity was generated during test network validator initialization
		require.Equal(t, len(objs)+1, int(resp.Pagination.Total))
		require.Subset(t,
			nullify.Fill(resp.DelegatedAmount),
			nullify.Fill(objs),
		)
	})
}
