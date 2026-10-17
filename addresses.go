package pxa

import "github.com/ethereum/go-ethereum/common"

// Testnet constants retain their historical deployment values.
// PXA mainnet uses bridge mirrors; direct PXT/PAT helpers are unavailable.
var (
	// PXTAddressTestnet is PXT (governance + utility) on PXA Testnet.
	PXTAddressTestnet = common.HexToAddress("0x254a549b711653263ff9EF22771565A50bC285c8")
	// PATAddressTestnet is PAT (activity reward) on PXA Testnet.
	PATAddressTestnet = common.HexToAddress("0x1333394DA43C81A0C91055F861A48A884FdF219B")
	// WPXTAddressTestnet is wPXT (bridge mirror) on PXA Testnet.
	WPXTAddressTestnet = common.HexToAddress("0xc60ef5398300437079F57B55fcaAa0A8B191BcC5")
	// WPATAddressTestnet is wPAT (bridge mirror) on PXA Testnet.
	WPATAddressTestnet = common.HexToAddress("0x0189a9C227C736D4dbaC3c010a54Cf55d7a323a5")
)

// Mainnet bridge-mirror addresses verified on 2026-10-18.
var (
	PXTAddressMainnet  = common.Address{}
	PATAddressMainnet  = common.Address{}
	WPXTAddressMainnet = common.HexToAddress("0x6848EB305D8Bc1ccB9a4b2Ab695F109fC25b80f3")
	WPATAddressMainnet = common.HexToAddress("0x32D2E7aBDecE846aDED04F71b509388B646A8741")
)
