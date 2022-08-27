[![Golang CI](https://github.com/rovergulf/eth-contracts-go/actions/workflows/checks.yml/badge.svg)](https://github.com/rovergulf/eth-contracts-go/actions/workflows/checks.yml)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# ETH Contracts Go

This repository contains Ethereum Smart Contracts bindings and some high-level API built on most popular interfaces

Smart contracts code is depending
on [OpenZeppelin contracts](https://github.com/OpenZeppelin/openzeppelin-contracts) used
in [eth-contracts](https://github.com/rovergulf/eth-contracts) repository.

## Supported interfaces

### Final

#### Access
- ERC173 - [EIP173](https://eips.ethereum.org/EIPS/eip-173) - Ownable

#### Introspection:
- ERC165 - [EIP165](https://eips.ethereum.org/EIPS/eip-165) - Standard Interface Detection
- ERC1820 - [EIP1820](https://eips.ethereum.org/EIPS/eip-1820) - Pseudo-introspection Registry Contract

#### Tokens:
- ERC20 - [EIP20](https://eips.ethereum.org/EIPS/eip-20) - Token Standard
- ERC721 - [EIP721](https://eips.ethereum.org/EIPS/eip-721) - Non-Fungible Tokens
- ERC777 - [EIP777](https://eips.ethereum.org/EIPS/eip-777) - Token Standard
- ERC1155 - [EIP1155](https://eips.ethereum.org/EIPS/eip-1155) - Multi-token Standard
- ERC4626 - [EIP4626](https://eips.ethereum.org/EIPS/eip-4626) - Tokenized Vault Standard

### Last call

#### TBD

## Usage

Use as dependency

```shell
go get github.com/rovergulf/eth-contracts-go
```

Install as CLI debug tool

```shell
go install github.com/rovergulf/eth-contracts-go

# if GOBIN specified
eth-contracts-go help
```
