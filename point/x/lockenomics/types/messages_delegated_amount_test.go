package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"point/testutil/sample"
)

func TestMsgCreateDelegatedAmount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateDelegatedAmount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateDelegatedAmount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateDelegatedAmount{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateDelegatedAmount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateDelegatedAmount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateDelegatedAmount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateDelegatedAmount{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteDelegatedAmount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteDelegatedAmount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteDelegatedAmount{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteDelegatedAmount{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
