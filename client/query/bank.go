package query

import (
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// bank_ParamsRPC returns the distribution params
func bankParamsRPC(q *Query) (*banktypes.QueryParamsResponse, error) {
	req := &banktypes.QueryParamsRequest{}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Params(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bank_BalanceRPC returns the balance of specified denom coins for a single account.
func bankBalanceRPC(q *Query, address string, denom string) (*banktypes.QueryBalanceResponse, error) {
	req := &banktypes.QueryBalanceRequest{Address: address, Denom: denom}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.Balance(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bank_AllBalancesRPC returns the balance of all coins for a single account.
func bankAllBalancesRPC(q *Query, address string) (*banktypes.QueryAllBalancesResponse, error) {
	req := &banktypes.QueryAllBalancesRequest{Address: address, Pagination: q.Options.Pagination}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.AllBalances(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bank_SupplyOfRPC returns the supply of all coins
func bankSupplyOfRPC(q *Query, denom string) (*banktypes.QuerySupplyOfResponse, error) {
	req := &banktypes.QuerySupplyOfRequest{Denom: denom}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.SupplyOf(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bank_TotalSupplyRPC returns the supply of all coins
func bankTotalSupplyRPC(q *Query) (*banktypes.QueryTotalSupplyResponse, error) {
	req := &banktypes.QueryTotalSupplyRequest{Pagination: q.Options.Pagination}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.TotalSupply(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bank_DenomMetadataRPC returns the metadata for given denom
func bankDenomMetadataRPC(q *Query, denom string) (*banktypes.QueryDenomMetadataResponse, error) {
	req := &banktypes.QueryDenomMetadataRequest{Denom: denom}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DenomMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// bank_DenomsMetadataRPC returns the metadata for all denoms
func bankDenomsMetadataRPC(q *Query) (*banktypes.QueryDenomsMetadataResponse, error) {
	req := &banktypes.QueryDenomsMetadataRequest{Pagination: q.Options.Pagination}
	queryClient := banktypes.NewQueryClient(q.Client)
	ctx, cancel := q.GetQueryContext()
	defer cancel()
	res, err := queryClient.DenomsMetadata(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
