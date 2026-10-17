# Go prerelease distribution

Source: https://github.com/tiling-pxa/pxa-go-sdk . Module path matches the public repository.

```sh
go get github.com/tiling-pxa/pxa-go-sdk@v0.1.0-rc.1
```

This is a prerelease. Use the committed go.mod/go.sum and toolchain for reproducible validation. Run `go test ./...` and `go vet ./...`. The included deployment ABI fixture makes the source independently testable. Verify public proxy installation and the tag commit after release; do not overwrite a published tag. No transaction or staking write is needed for verification.
