[![Golang CI](https://github.com/rovergulf/eth-contracts-go/actions/workflows/checks.yml/badge.svg)](https://github.com/rovergulf/eth-contracts-go/actions/workflows/checks.yml)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)



# ETH Contracts Go

This repository contains Ethereum Smart Contracts bindings using Golang generated ABI code

Smart contracts code is depending
on [OpenZeppelin contracts](https://github.com/OpenZeppelin/openzeppelin-contracts) used
in [eth-contracts](https://github.com/rovergulf/eth-contracts) repository.

### Supported final interfaces

- ERC20 - [EIP20](https://eips.ethereum.org/EIPS/eip-20)
- ERC165 - [EIP165](https://eips.ethereum.org/EIPS/eip-165)
- ERC721 - [EIP721](https://eips.ethereum.org/EIPS/eip-721)
- ERC777 - [EIP777](https://eips.ethereum.org/EIPS/eip-777)
- ERC1155 - [EIP1155](https://eips.ethereum.org/EIPS/eip-1155)

### Use as dependency

```shell
go get github.com/rovergulf/eth-contracts-go
```

### Install as CLI debug tool

```shell
go install github.com/rovergulf/eth-contracts-go

# if GOBIN specified
eth-contracts-go help
```
