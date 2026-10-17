package pxa

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// ERC20ABIJSON is the standard ERC-20 ABI as a JSON string, mirroring
// @pxachain/sdk src/abis/erc20.ts. PXT and PAT use this verbatim.
const ERC20ABIJSON = `[
  {"type":"function","name":"name","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"string"}]},
  {"type":"function","name":"symbol","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"string"}]},
  {"type":"function","name":"decimals","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"uint8"}]},
  {"type":"function","name":"totalSupply","stateMutability":"view","inputs":[],"outputs":[{"name":"","type":"uint256"}]},
  {"type":"function","name":"balanceOf","stateMutability":"view","inputs":[{"name":"owner","type":"address"}],"outputs":[{"name":"","type":"uint256"}]},
  {"type":"function","name":"allowance","stateMutability":"view","inputs":[{"name":"owner","type":"address"},{"name":"spender","type":"address"}],"outputs":[{"name":"","type":"uint256"}]},
  {"type":"function","name":"transfer","stateMutability":"nonpayable","inputs":[{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"outputs":[{"name":"","type":"bool"}]},
  {"type":"function","name":"approve","stateMutability":"nonpayable","inputs":[{"name":"spender","type":"address"},{"name":"value","type":"uint256"}],"outputs":[{"name":"","type":"bool"}]},
  {"type":"function","name":"transferFrom","stateMutability":"nonpayable","inputs":[{"name":"from","type":"address"},{"name":"to","type":"address"},{"name":"value","type":"uint256"}],"outputs":[{"name":"","type":"bool"}]},
  {"type":"event","name":"Transfer","anonymous":false,"inputs":[{"name":"from","type":"address","indexed":true},{"name":"to","type":"address","indexed":true},{"name":"value","type":"uint256","indexed":false}]},
  {"type":"event","name":"Approval","anonymous":false,"inputs":[{"name":"owner","type":"address","indexed":true},{"name":"spender","type":"address","indexed":true},{"name":"value","type":"uint256","indexed":false}]}
]`

// WrappedTokenABIJSON is the minimal ERC-20 surface of the bridge mirrors.
// These contracts are not WETH and do not expose deposit/withdraw.
const WrappedTokenABIJSON = ERC20ABIJSON

// ERC20ABI is the parsed standard ERC-20 ABI, ready for binding.
func ERC20ABI() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(ERC20ABIJSON))
}

// WrappedTokenABI is the parsed bridge-mirror ERC-20
// ABI, ready for binding.
func WrappedTokenABI() (abi.ABI, error) {
	return abi.JSON(strings.NewReader(WrappedTokenABIJSON))
}
