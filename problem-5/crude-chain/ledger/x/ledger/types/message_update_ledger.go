package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateLedger{}

func NewMsgUpdateLedger(creator string, title string, body string, cost uint64, id uint64) *MsgUpdateLedger {
	return &MsgUpdateLedger{
		Creator: creator,
		Title:   title,
		Body:    body,
		Cost:    cost,
		Id:      id,
	}
}

func (msg *MsgUpdateLedger) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
