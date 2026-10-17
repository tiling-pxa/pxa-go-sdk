package pxa_test

import (
	"fmt"

	pxa "github.com/tiling-pxa/pxa-go-sdk"
)

// ExampleChainByID shows resolving a PXA chain by its id.
func ExampleChainByID() {
	c, ok := pxa.ChainByID(7840)
	fmt.Println(ok, c.Name, c.NativeCurrency.Symbol)
	// Output: true PXA Testnet tPXA
}

// ExampleGetPXTContract shows resolving the PXT token address for a
// chain without dialing a node (pass a real backend in production).
func ExampleGetPXTContract() {
	c, err := pxa.GetPXTContract(pxa.PXATestnet, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(c.Address.Hex())
	// Output: 0x254a549b711653263ff9EF22771565A50bC285c8
}

// Example_dialAndReadBalance is the canonical end-to-end shape: dial a
// PXA chain, bind PXT, and call balanceOf. It is illustrative only
// (not run as a test) because it requires a live RPC endpoint.
func Example_dialAndReadBalance() {
	// ctx := context.Background()
	// client, err := pxa.Dial(ctx, pxa.PXATestnet)
	// if err != nil { log.Fatal(err) }
	// defer client.Close()
	//
	// pxt, err := pxa.GetPXTContract(pxa.PXATestnet, client)
	// if err != nil { log.Fatal(err) }
	//
	// var out []interface{}
	// err = pxt.Contract.Call(&bind.CallOpts{Context: ctx}, &out, "balanceOf", addr)
	// if err != nil { log.Fatal(err) }
	// fmt.Println(out[0])
}
