package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"point/testutil/sample"
)

func TestMsgCreateLock_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateLock
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateLock{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateLock{
				Creator: sample.AccAddress(),
			},
		}, {
			name: "delegation is too long",
			msg: MsgCreateLock{
				Creator: sample.AccAddress(),
				Lenght:  MaxDelegationLength + 1,
			},
			err: ErrTooLongDelegation,
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
