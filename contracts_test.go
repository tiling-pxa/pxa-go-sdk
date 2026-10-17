package pxa

import (
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

// Mirrors @pxachain/sdk test/contracts.test.ts. A nil backend is fine
// here: NewBoundContract stores the backend without calling it, and we
// only assert on address + ABI resolution.

func TestGetPXTContract(t *testing.T) {
	c, err := GetPXTContract(PXATestnet, nil)
	if err != nil {
		t.Fatalf("testnet: unexpected error: %v", err)
	}
	if c.Address != PXTAddressTestnet {
		t.Errorf("PXT testnet address = %s, want %s", c.Address, PXTAddressTestnet)
	}

	if _, err := GetPXTContract(PXAMainnet, nil); err == nil || !strings.Contains(err.Error(), "no direct PXA mainnet contract") {
		t.Fatalf("mainnet direct helper must reject: %v", err)
	}

	bad := Chain{ID: 9999}
	if _, err := GetPXTContract(bad, nil); err == nil {
		t.Error("expected error for unsupported chain id 9999")
	} else if !strings.Contains(err.Error(), "not configured for chain id 9999") {
		t.Errorf("error = %q, want mention of chain id 9999", err.Error())
	}
}

func TestGetPATContract(t *testing.T) {
	c, err := GetPATContract(PXATestnet, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Address != PATAddressTestnet {
		t.Errorf("PAT testnet address = %s, want %s", c.Address, PATAddressTestnet)
	}

	bad := Chain{ID: 1}
	if _, err := GetPATContract(bad, nil); err == nil {
		t.Error("expected error for unsupported chain id 1")
	} else if !strings.Contains(err.Error(), "PAT not configured for chain id 1") {
		t.Errorf("error = %q", err.Error())
	}
}

func TestGetWrappedContracts(t *testing.T) {
	w1, err := GetWPXTContract(PXATestnet, nil)
	if err != nil {
		t.Fatalf("wPXT: %v", err)
	}
	if w1.Address != WPXTAddressTestnet {
		t.Errorf("wPXT testnet address = %s, want %s", w1.Address, WPXTAddressTestnet)
	}
	// Bridge mirrors must not expose WETH methods.
	if _, ok := w1.ABI.Methods["deposit"]; ok {
		t.Error("wPXT ABI falsely exposes deposit")
	}
	if _, ok := w1.ABI.Methods["withdraw"]; ok {
		t.Error("wPXT ABI falsely exposes withdraw")
	}

	w2, err := GetWPATContract(PXATestnet, nil)
	if err != nil {
		t.Fatalf("wPAT: %v", err)
	}
	if w2.Address != WPATAddressTestnet {
		t.Errorf("wPAT testnet address = %s, want %s", w2.Address, WPATAddressTestnet)
	}
}

func TestABIsParse(t *testing.T) {
	erc20, err := ERC20ABI()
	if err != nil {
		t.Fatalf("ERC20ABI parse: %v", err)
	}
	for _, m := range []string{"name", "symbol", "decimals", "totalSupply", "balanceOf", "allowance", "transfer", "approve", "transferFrom"} {
		if _, ok := erc20.Methods[m]; !ok {
			t.Errorf("ERC20 ABI missing method %q", m)
		}
	}
	if _, ok := erc20.Events["Transfer"]; !ok {
		t.Error("ERC20 ABI missing Transfer event")
	}

	wrapped, err := WrappedTokenABI()
	if err != nil {
		t.Fatalf("WrappedTokenABI parse: %v", err)
	}
	if _, ok := wrapped.Events["Deposit"]; ok {
		t.Error("wrapped ABI falsely exposes Deposit")
	}
	if _, ok := wrapped.Events["Withdrawal"]; ok {
		t.Error("wrapped ABI falsely exposes Withdrawal")
	}
}

// TestTestnetAddressesAreChecksummed guards the EIP-55 checksum form of
// the testnet addresses against accidental lowercase edits — they must
// round-trip through common.HexToAddress to the documented mixed-case.
func TestTestnetAddressesAreChecksummed(t *testing.T) {
	cases := map[string]common.Address{
		"0x254a549b711653263ff9EF22771565A50bC285c8": PXTAddressTestnet,
		"0x1333394DA43C81A0C91055F861A48A884FdF219B": PATAddressTestnet,
		"0xc60ef5398300437079F57B55fcaAa0A8B191BcC5": WPXTAddressTestnet,
		"0x0189a9C227C736D4dbaC3c010a54Cf55d7a323a5": WPATAddressTestnet,
	}
	for want, got := range cases {
		if got.Hex() != want {
			t.Errorf("address %s checksums to %s", want, got.Hex())
		}
	}
}
