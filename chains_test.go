package pxa

import "testing"

// These tests mirror @pxachain/sdk test/chains.test.ts so the two SDKs
// stay byte-equal on the canonical chain parameters.

func TestPXATestnet(t *testing.T) {
	if PXATestnet.ID != 7840 {
		t.Errorf("testnet id = %d, want 7840", PXATestnet.ID)
	}
	if PXATestnet.RPCURL != "https://testnet-rpc.pxachain.com" {
		t.Errorf("testnet RPC = %q", PXATestnet.RPCURL)
	}
	if PXATestnet.Explorer.URL != "https://testnet-explorer.pxachain.com" {
		t.Errorf("testnet explorer = %q", PXATestnet.Explorer.URL)
	}
	if PXATestnet.NativeCurrency.Symbol != "tPXA" {
		t.Errorf("testnet native symbol = %q, want tPXA", PXATestnet.NativeCurrency.Symbol)
	}
	if PXATestnet.NativeCurrency.Decimals != 18 {
		t.Errorf("testnet decimals = %d, want 18", PXATestnet.NativeCurrency.Decimals)
	}
	if !PXATestnet.Testnet {
		t.Error("testnet flag should be true")
	}
}

func TestPXAMainnet(t *testing.T) {
	if PXAMainnet.ID != 784 {
		t.Errorf("mainnet id = %d, want 784", PXAMainnet.ID)
	}
	if PXAMainnet.RPCURL != "https://rpc.pxachain.com" {
		t.Errorf("mainnet RPC = %q", PXAMainnet.RPCURL)
	}
	if PXAMainnet.Explorer.URL != "https://explorer.pxachain.com" {
		t.Errorf("mainnet explorer = %q", PXAMainnet.Explorer.URL)
	}
	if PXAMainnet.NativeCurrency.Symbol != "PXA" {
		t.Errorf("mainnet native symbol = %q, want PXA", PXAMainnet.NativeCurrency.Symbol)
	}
	if PXAMainnet.Testnet {
		t.Error("mainnet flag should be false")
	}
}

func TestPXAChains(t *testing.T) {
	if len(PXAChains) != 2 {
		t.Fatalf("len(PXAChains) = %d, want 2", len(PXAChains))
	}
	if PXAChains[0].ID != PXATestnet.ID {
		t.Error("PXAChains[0] should be testnet")
	}
	if PXAChains[1].ID != PXAMainnet.ID {
		t.Error("PXAChains[1] should be mainnet")
	}
}

func TestChainByID(t *testing.T) {
	if c, ok := ChainByID(7840); !ok || c.ID != 7840 {
		t.Errorf("ChainByID(7840) = %v, %v", c.ID, ok)
	}
	if c, ok := ChainByID(784); !ok || c.ID != 784 {
		t.Errorf("ChainByID(784) = %v, %v", c.ID, ok)
	}
	if _, ok := ChainByID(1); ok {
		t.Error("ChainByID(1) should not resolve to a PXA chain")
	}
}

func TestEIP3091URLs(t *testing.T) {
	if got := PXATestnet.TxURL("0xabc"); got != "https://testnet-explorer.pxachain.com/tx/0xabc" {
		t.Errorf("TxURL = %q", got)
	}
	if got := PXATestnet.BlockURL("123"); got != "https://testnet-explorer.pxachain.com/block/123" {
		t.Errorf("BlockURL = %q", got)
	}
}
