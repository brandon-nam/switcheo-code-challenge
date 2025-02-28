package keeper

import (
	"context"
	"fmt"

	"ledger/x/ledger/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteLedger(goCtx context.Context, msg *types.MsgDeleteLedger) (*types.MsgDeleteLedgerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetLedger(ctx, msg.Id)
    if !found {
        return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
    }
    if msg.Creator != val.Creator {
        return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
    }
    k.RemoveLedger(ctx, msg.Id)
    return &types.MsgDeleteLedgerResponse{}, nil
}
