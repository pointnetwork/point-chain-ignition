package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDelegatedAmount = "create_delegated_amount"
	TypeMsgUpdateDelegatedAmount = "update_delegated_amount"
	TypeMsgDeleteDelegatedAmount = "delete_delegated_amount"
)

var _ sdk.Msg = &MsgCreateDelegatedAmount{}

func NewMsgCreateDelegatedAmount(
	creator string,
	index string,
	delegator string,
	validator string,
	amount uint64,

) *MsgCreateDelegatedAmount {
	return &MsgCreateDelegatedAmount{
		Creator:   creator,
		Index:     index,
		Delegator: delegator,
		Validator: validator,
		Amount:    amount,
	}
}

func (msg *MsgCreateDelegatedAmount) Route() string {
	return RouterKey
}

func (msg *MsgCreateDelegatedAmount) Type() string {
	return TypeMsgCreateDelegatedAmount
}

func (msg *MsgCreateDelegatedAmount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDelegatedAmount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDelegatedAmount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDelegatedAmount{}

func NewMsgUpdateDelegatedAmount(
	creator string,
	index string,
	delegator string,
	validator string,
	amount uint64,

) *MsgUpdateDelegatedAmount {
	return &MsgUpdateDelegatedAmount{
		Creator:   creator,
		Index:     index,
		Delegator: delegator,
		Validator: validator,
		Amount:    amount,
	}
}

func (msg *MsgUpdateDelegatedAmount) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDelegatedAmount) Type() string {
	return TypeMsgUpdateDelegatedAmount
}

func (msg *MsgUpdateDelegatedAmount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDelegatedAmount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDelegatedAmount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDelegatedAmount{}

func NewMsgDeleteDelegatedAmount(
	creator string,
	index string,

) *MsgDeleteDelegatedAmount {
	return &MsgDeleteDelegatedAmount{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteDelegatedAmount) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDelegatedAmount) Type() string {
	return TypeMsgDeleteDelegatedAmount
}

func (msg *MsgDeleteDelegatedAmount) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDelegatedAmount) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDelegatedAmount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
