package main

import (
	"context"
	"fmt"
)

// Fetches price
type PriceFetcher interface {
	FetchPrice(context.Context, string) (float64, error)
}

// Implements PriceFeter interface
type priceFetcher struct{}

func (s *priceFetcher) FetchPrice(ctx context.Context, ticker string) (float64, error) {
	return MockPriceFetcher(ctx, ticker)
}

var priceMock = map[string]float64{
	"BTC": 20_000.0,
	"ETH": 200.0,
}

// Replace this with Chainlink orcal to get prices
func MockPriceFetcher(ctx context.Context, ticker string) (float64, error) {
	price, ok := priceMock[ticker]

	if !ok {
		return price, fmt.Errorf("The given ticker (%s) is not supported", ticker)
	}

	return price, nil
}
