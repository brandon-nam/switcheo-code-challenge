package keeper

import (
	"context"

	"ledger/x/ledger/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowLedger(goCtx context.Context, req *types.QueryShowLedgerRequest) (*types.QueryShowLedgerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetLedger(ctx, req.Id)

	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Ledger of such ID not found!")
	}

	return &types.QueryShowLedgerResponse{Ledger: &val}, nil
}
