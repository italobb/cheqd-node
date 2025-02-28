package keeper

import (
	"github.com/cheqd/cheqd-node/x/cheqd/types/v1"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case v1.QueryGetDid:
			return getDid(ctx, path[1], k, legacyQuerierCdc)

		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", v1.ModuleName, path[0])
		}

		return res, err
	}
}
