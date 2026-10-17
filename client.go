package pxa

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

// Dial connects to a PXA chain's default HTTPS RPC endpoint and returns
// a go-ethereum *ethclient.Client. It is a thin convenience over
// ethclient.Dial(chain.RPCURL) — pass the returned client to the
// Get*Contract helpers as the backend.
//
// For a custom RPC URL (your own node, a private gateway), call
// ethclient.Dial directly with that URL instead.
func Dial(ctx context.Context, chain Chain) (*ethclient.Client, error) {
	client, err := ethclient.DialContext(ctx, chain.RPCURL)
	if err != nil {
		return nil, fmt.Errorf("pxa-go-sdk: dial %s (%s): %w", chain.Name, chain.RPCURL, err)
	}
	return client, nil
}
