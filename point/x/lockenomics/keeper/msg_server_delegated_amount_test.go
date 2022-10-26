package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "point/testutil/keeper"
	"point/x/lockenomics/keeper"
	"point/x/lockenomics/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestDelegatedAmountMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.LockenomicsKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateDelegatedAmount{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateDelegatedAmount(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetDelegatedAmount(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestDelegatedAmountMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateDelegatedAmount
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateDelegatedAmount{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateDelegatedAmount{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateDelegatedAmount{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LockenomicsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateDelegatedAmount{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateDelegatedAmount(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateDelegatedAmount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetDelegatedAmount(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestDelegatedAmountMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteDelegatedAmount
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteDelegatedAmount{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteDelegatedAmount{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteDelegatedAmount{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.LockenomicsKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateDelegatedAmount(wctx, &types.MsgCreateDelegatedAmount{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteDelegatedAmount(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetDelegatedAmount(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
