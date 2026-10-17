// Package pxa is the Go SDK for PXA Chain (testnet + mainnet).
//
// It mirrors @pxachain/sdk (TypeScript): the same chain definitions,
// the same four official-token addresses, the same ABIs, and the same
// per-token contract helpers. Values are kept byte-equal with the
// canonical sources via the drift test in chains_test.go.
//
// PXA is EVM-compatible (a BSC/Geth fork running Parlia consensus), so
// this SDK is a thin convenience layer over go-ethereum: it does not
// wrap the RPC client, it just hands you the chain constants, addresses,
// ABIs, and bound contracts so you skip the boilerplate.
package pxa

import "github.com/ethereum/go-ethereum/common"

// NativeCurrency describes a chain's gas token.
type NativeCurrency struct {
	Name     string
	Symbol   string
	Decimals uint8
}

// Explorer is an EIP-3091-compatible block explorer entry.
type Explorer struct {
	Name string
	URL  string
}

// Chain is a typed PXA network definition. It mirrors the viem
// `defineChain` shape exported by @pxachain/sdk.
type Chain struct {
	// ID is the EIP-155 chainId in decimal.
	ID int64
	// Name is the human-readable chain name.
	Name string
	// NativeCurrency is the gas token.
	NativeCurrency NativeCurrency
	// RPCURL is the default HTTPS JSON-RPC endpoint.
	RPCURL string
	// WSURL is the default WebSocket JSON-RPC endpoint.
	WSURL string
	// Explorer is the default block explorer.
	Explorer Explorer
	// Testnet flags the network as a test network.
	Testnet bool
}

// PXATestnet is the PXA Testnet (chainId 7840).
//
// Native currency is tPXA to keep wallets from confusing testnet PXA
// with the mainnet token. Mirrors chains.ts::pxaTestnet.
var PXATestnet = Chain{
	ID:             7840,
	Name:           "PXA Testnet",
	NativeCurrency: NativeCurrency{Name: "tPXA", Symbol: "tPXA", Decimals: 18},
	RPCURL:         "https://testnet-rpc.pxachain.com",
	WSURL:          "wss://testnet-wss.pxachain.com",
	Explorer: Explorer{
		Name: "PXA Testnet Explorer",
		URL:  "https://testnet-explorer.pxachain.com",
	},
	Testnet: true,
}

// PXAMainnet is the PXA Chain mainnet (chainId 784).
//
// The "launched" flag lives only in the docs site (it gates the
// "Coming soon" badge); the SDK exposes the chain definition
// unconditionally so apps can type-check before launch. Mirrors
// chains.ts::pxaMainnet.
var PXAMainnet = Chain{
	ID:             784,
	Name:           "PXA Chain",
	NativeCurrency: NativeCurrency{Name: "PXA", Symbol: "PXA", Decimals: 18},
	RPCURL:         "https://rpc.pxachain.com",
	WSURL:          "wss://rpc.pxachain.com/ws",
	Explorer: Explorer{
		Name: "PXA Chain Explorer",
		URL:  "https://explorer.pxachain.com",
	},
	Testnet: false,
}

// PXAChains is the ordered list of PXA networks: [testnet, mainnet].
// Mirrors chains.ts::pxaChains.
var PXAChains = []Chain{PXATestnet, PXAMainnet}

// ChainByID returns the PXA chain definition for the given chainId, and
// ok=false if the id is not a PXA network.
func ChainByID(id int64) (Chain, bool) {
	for _, c := range PXAChains {
		if c.ID == id {
			return c, true
		}
	}
	return Chain{}, false
}

// EIP-3091 deep-link helpers. They mirror the explorer URL patterns
// documented at /network-info, so callers don't hand-assemble URLs.

// BlockURL returns the explorer URL for a block (by height or hash).
func (c Chain) BlockURL(heightOrHash string) string {
	return c.Explorer.URL + "/block/" + heightOrHash
}

// TxURL returns the explorer URL for a transaction hash.
func (c Chain) TxURL(hash string) string {
	return c.Explorer.URL + "/tx/" + hash
}

// AddressURL returns the explorer URL for an account address.
func (c Chain) AddressURL(addr common.Address) string {
	return c.Explorer.URL + "/address/" + addr.Hex()
}

// TokenURL returns the explorer URL for a token contract address.
func (c Chain) TokenURL(addr common.Address) string {
	return c.Explorer.URL + "/token/" + addr.Hex()
}
