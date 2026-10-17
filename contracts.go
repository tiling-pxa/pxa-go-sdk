package pxa

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// Per-token contract helpers. Each helper:
//  1. Picks the right address from the supplied chain (testnet vs mainnet).
//  2. Returns an error on any other chain id rather than silently
//     binding to the zero address.
//  3. Binds the token ABI to the supplied backend via go-ethereum's
//     bind.NewBoundContract.
//
// This mirrors @pxachain/sdk src/contracts.ts. The backend is any
// type implementing bind.ContractBackend — *ethclient.Client is the
// usual choice (see the run-a-node / SDK docs for wiring).

// TokenContract bundles a token's address, parsed ABI, and bound
// contract so callers can introspect or invoke without re-deriving any
// of the three.
type TokenContract struct {
	Address  common.Address
	ABI      abi.ABI
	Contract *bind.BoundContract
}

// pickAddress returns the per-chain address for a token, or an error if
// the chain is not a PXA network. Mirrors contracts.ts::pickAddress.
func pickAddress(chain Chain, testnetAddr, mainnetAddr common.Address, symbol string) (common.Address, error) {
	switch chain.ID {
	case PXATestnet.ID:
		return testnetAddr, nil
	case PXAMainnet.ID:
		if mainnetAddr == (common.Address{}) {
			return common.Address{}, fmt.Errorf("pxa-go-sdk: %s has no direct PXA mainnet contract; use the backed wrapped token helper", symbol)
		}
		return mainnetAddr, nil
	default:
		return common.Address{}, fmt.Errorf(
			"pxa-go-sdk: %s not configured for chain id %d; supported chains are 7840 (testnet) and 784 (mainnet)",
			symbol, chain.ID,
		)
	}
}

// bindToken parses the ABI, resolves the address, and binds the
// contract to the backend.
func bindToken(chain Chain, testnetAddr, mainnetAddr common.Address, symbol, abiJSON string, backend bind.ContractBackend) (*TokenContract, error) {
	address, err := pickAddress(chain, testnetAddr, mainnetAddr, symbol)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, fmt.Errorf("pxa-go-sdk: parse %s ABI: %w", symbol, err)
	}
	return &TokenContract{
		Address:  address,
		ABI:      parsed,
		Contract: bind.NewBoundContract(address, parsed, backend, backend, backend),
	}, nil
}

// GetPXTContract binds PXT (governance + utility, ERC-20) for the chain.
func GetPXTContract(chain Chain, backend bind.ContractBackend) (*TokenContract, error) {
	return bindToken(chain, PXTAddressTestnet, PXTAddressMainnet, "PXT", ERC20ABIJSON, backend)
}

// GetPATContract binds PAT (activity reward, ERC-20) for the chain.
func GetPATContract(chain Chain, backend bind.ContractBackend) (*TokenContract, error) {
	return bindToken(chain, PATAddressTestnet, PATAddressMainnet, "PAT", ERC20ABIJSON, backend)
}

// GetWPXTContract binds wPXT (bridge mirror, minimal ERC-20).
func GetWPXTContract(chain Chain, backend bind.ContractBackend) (*TokenContract, error) {
	return bindToken(chain, WPXTAddressTestnet, WPXTAddressMainnet, "wPXT", WrappedTokenABIJSON, backend)
}

// GetWPATContract binds wPAT (bridge mirror, minimal ERC-20).
func GetWPATContract(chain Chain, backend bind.ContractBackend) (*TokenContract, error) {
	return bindToken(chain, WPATAddressTestnet, WPATAddressMainnet, "wPAT", WrappedTokenABIJSON, backend)
}
