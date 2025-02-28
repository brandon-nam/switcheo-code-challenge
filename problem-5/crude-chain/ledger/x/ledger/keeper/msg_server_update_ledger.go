package keeper

import (
	"context"
	"fmt"

	"ledger/x/ledger/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateLedger(goCtx context.Context, msg *types.MsgUpdateLedger) (*types.MsgUpdateLedgerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var ledger = types.Ledger{
		Creator: msg.Creator,
		Id:      msg.Id,
		Title:   msg.Title,
		Body:    msg.Body,
	}
	val, found := k.GetLedger(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetLedger(ctx, ledger)
	return &types.MsgUpdateLedgerResponse{}, nil
}
