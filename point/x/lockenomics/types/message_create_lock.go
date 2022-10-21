package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateLock = "create_lock"

var _ sdk.Msg = &MsgCreateLock{}

func NewMsgCreateLock(creator string, lenght int32, validator string) *MsgCreateLock {
	return &MsgCreateLock{
		Creator:   creator,
		Lenght:    lenght,
		Validator: validator,
	}
}

func (msg *MsgCreateLock) Route() string {
	return RouterKey
}

func (msg *MsgCreateLock) Type() string {
	return TypeMsgCreateLock
}

func (msg *MsgCreateLock) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateLock) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateLock) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
