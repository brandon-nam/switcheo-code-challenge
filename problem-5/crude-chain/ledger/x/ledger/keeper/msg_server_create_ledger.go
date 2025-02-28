package keeper

import (
	"context"

	"ledger/x/ledger/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateLedger(goCtx context.Context, msg *types.MsgCreateLedger) (*types.MsgCreateLedgerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	newLedger := types.Ledger{
		Creator: msg.Creator,
		Cost:    msg.Cost,
		Title:   msg.Title,
		Body:    msg.Body,
	}

	id := k.AppendLedger(ctx, newLedger)

	return &types.MsgCreateLedgerResponse{Id: id}, nil
}
