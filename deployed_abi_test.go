package pxa

import (
	"os"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func TestDeployedABICompatibility(t *testing.T) {
	data, err := os.ReadFile("testdata/wrapped-bridge-mainnet-abi.json")
	if err != nil {
		t.Fatal(err)
	}
	deployed, err := abi.JSON(strings.NewReader(string(data)))
	if err != nil {
		t.Fatal(err)
	}
	exported, err := WrappedTokenABI()
	if err != nil {
		t.Fatal(err)
	}
	for name, method := range exported.Methods {
		target, ok := deployed.Methods[name]
		if !ok || target.Sig != method.Sig || target.StateMutability != method.StateMutability || len(target.Outputs) != len(method.Outputs) {
			t.Fatalf("method mismatch: %s", name)
		}
		for i, output := range method.Outputs {
			if target.Outputs[i].Type.String() != output.Type.String() {
				t.Fatalf("output mismatch: %s", name)
			}
		}
	}
	for name, event := range exported.Events {
		target, ok := deployed.Events[name]
		if !ok || target.Sig != event.Sig {
			t.Fatalf("event mismatch: %s", name)
		}
		for i, input := range event.Inputs {
			if target.Inputs[i].Indexed != input.Indexed {
				t.Fatalf("indexed mismatch: %s", name)
			}
		}
	}
	pxt, err := GetWPXTContract(PXAMainnet, nil)
	if err != nil {
		t.Fatal(err)
	}
	pat, err := GetWPATContract(PXAMainnet, nil)
	if err != nil {
		t.Fatal(err)
	}
	if strings.ToLower(pxt.Address.Hex()) != "0x6848eb305d8bc1ccb9a4b2ab695f109fc25b80f3" || strings.ToLower(pat.Address.Hex()) != "0x32d2e7abdece846aded04f71b509388b646a8741" {
		t.Fatal("mainnet address mismatch")
	}
	if _, err := GetPATContract(PXAMainnet, nil); err == nil {
		t.Fatal("direct PAT placeholder must reject")
	}
}
