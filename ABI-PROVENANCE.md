# Bridge mirror ABI evidence

The fixture is the public ABI of `contracts/bridge/tokens/WrappedBridgeToken.sol`, extracted from the archived production compiler output (Solidity 0.8.24+commit.e11b9ed9). Archive SHA256: `79487bc8f4bc43ba82c4cbb65697bc2bea5b848d11dec4eb8630024c02188ef3`.

On 2026-10-18, public JSON-RPC at PXA block 445265, hash `0x8c2689cd0da4971b6837271c3b202b945d11679c55a47d1c64d300194677fc55`, returned both mirror runtime codes. Each matched the archived runtime after masking only its seven compiler-declared 32-byte immutable references. This proves the checked code portion, not a byte-for-byte constructor/immutable configuration audit or independent security audit.

- PXT mirror `0x6848eb305d8bc1ccb9a4b2ab695f109fc25b80f3`: runtime SHA256 `b6fbfefbadf294dcbef5a9942980c0a5bade7d66d251ff44a6b2758e0e9d39c2`.
- PAT mirror `0x32d2e7abdece846aded04f71b509388b646a8741`: runtime SHA256 `cc943e64ffd2e76dfdab5212b4fd24f73f9034bd157562b8e06d3fa4b2932673`.

The SDK exports only matching standard ERC-20 entries from this ABI. WETH deposit/withdraw functions and Deposit/Withdrawal events do not exist in the fixture and must not be exposed by mirror helpers. Administrative bridge methods are intentionally outside this minimal SDK surface. No testnet runtime equivalence or transaction execution was established.

Go `testdata/` and Python `tests/fixtures/` contain byte-identical copies so standalone source distributions can run the ABI regression without the monorepo.
