# pxa-go-sdk — release candidate source

Go module `github.com/tiling-pxa/pxa-go-sdk`, using the committed go.mod / go.sum and toolchain. Public source and prerelease tags use the `tiling-pxa/pxa-go-sdk` repository. Release candidate: `v0.1.0-rc.2`.

## Validate the source

```sh
GOTOOLCHAIN=go1.25.11 go test ./...
GOTOOLCHAIN=go1.25.11 go vet ./...
```

For a clean consumer, use a local `replace github.com/tiling-pxa/pxa-go-sdk => /absolute/path/to/candidate` and test the actual package. Public module installation must be checked separately after repository and tag publication.

`Dial(ctx, PXAMainnet)` returns an ethclient; compare its `ChainID` to `PXAMainnet.ID` before reads. `GetWPXTContract(PXAMainnet, client)` and `GetWPATContract` bind the two mainnet mirrors. No signer or private key is required for reads.

## Deployment scope

PXA mainnet chain ID is 784. The mirror addresses are:

- PXT mirror: `0x6848EB305D8Bc1ccB9a4b2Ab695F109fC25b80f3` (on-chain symbol `PXT`).
- PAT mirror: `0x32D2E7aBDecE846aDED04F71b509388B646A8741` (on-chain symbol `PAT`).

The `wPXT`/`wPAT` helper names describe bridge roles, not current symbols. The ABI is a **minimal ERC-20 subset**, not the full administrative contract ABI. These mirrors do not have WETH-style `deposit`, `withdraw`, `Deposit` or `Withdrawal`. Direct PXT/PAT mainnet helpers reject instead of binding a zero address; use the mirror helpers on PXA, not BSC token addresses.

ABI provenance: archived Solidity 0.8.24 deployment artifact, compared to both mainnet runtime codes at block 445265 on 18 October 2026, excluding the seven compiler-declared immutable slots. See `ABI-PROVENANCE.md` for exact evidence and limits. Tests compare exported signatures against that independent ABI fixture. Testnet constants are historical and were not live-validated in this release.

A successful read does not enable transaction flows. Website staking writes remain closed. Never send wallet secrets in a diagnostic report.

## Distribution

Public source is maintained at https://github.com/tiling-pxa/pxa-go-sdk. GitHub prereleases distribute reviewed candidates. npm/PyPI publication is not available yet; do not install similarly named registry packages. Website staking writes remain closed. See `RELEASE.md` for installation and verification.
