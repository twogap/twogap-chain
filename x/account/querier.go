package account

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1:], req, keeper)
		case QueryWhois:
			return queryWhois(ctx, path[1:], req, keeper)
		case QueryNames:
			return queryNames(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown account query endpoint")
		}
	}
}

func queryResolve(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	name := path[0]

	value := keeper.ResolveName(ctx, name)

	if value == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not resolve name")
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResResolve{value})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

type QueryResResolve struct {
	Value string `json:"value"`
}

func (r QueryResResolve) String() string {
	return r.Value
}

func queryWhois(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	name := path[0]

	whois := keeper.GetWhois(ctx, name)

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, whois)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func (w Whois) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Value: %s
Price: %s`, w.Owner, w.Value, w.Price))
}

func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) (res []byte, err sdk.Error) {
	var namesList QueryResNames

	iterator := keeper.GetNamesIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		name := string(iterator.Key())
		namesList = append(namesList, name)
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, namesList)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

type QueryResNames []string

func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}
