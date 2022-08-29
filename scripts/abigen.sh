#!/usr/bin/env bash

set -e

source ./scripts/util.sh

mkdir -p tmp/abigen
mkdir -p tmp/bytecode
jsonnet hack/jsonnet/abi.jsonnet -m tmp/abigen
jsonnet hack/jsonnet/bin.jsonnet -m tmp/bytecode

for bin in tmp/bytecode/*; do
  echo $bin
  value=$(cat "$bin" | tr -d '"')
  echo $value > $bin
done

abigen --abi=./tmp/abigen/erc165 --pkg erc165 --out=abis/introspection/erc165/erc165.go
abigen --abi=./tmp/abigen/erc1820Registry --pkg erc1820 --out=abis/introspection/erc1820/erc1820Registry.go

abigen --abi=./tmp/abigen/ownable --pkg ownable --out=abis/access/ownable/ownable.go
abigen --abi=./tmp/abigen/erc173 --pkg erc173 --out=abis/interfaces/erc173/erc173.go
abigen --abi=./tmp/abigen/erc5313 --pkg erc5313 --out=abis/interfaces/erc5313/erc5313.go

abigen --abi=./tmp/abigen/erc4626 --pkg erc4626 --out=abis/token/erc4626/erc4626.go
abigen --abi=./tmp/abigen/erc20 --bin=./tmp/bytecode/erc20 --pkg erc20 --out=abis/token/erc20/erc20.go
abigen --abi=./tmp/abigen/erc721 --bin=./tmp/bytecode/erc721 --pkg erc721 --out=abis/token/erc721/erc721.go
abigen --abi=./tmp/abigen/erc721Enumerable --pkg erc721Enumerable --out=abis/token/erc721/extensions/erc721Enumerable/erc721Enumerable.go
abigen --abi=./tmp/abigen/erc721Metadata --pkg erc721Metadata --out=abis/token/erc721/extensions/erc721Metadata/erc721Metadata.go
abigen --abi=./tmp/abigen/erc777 --bin=./tmp/bytecode/erc777 --pkg erc777 --out=abis/token/erc777/erc777.go
abigen --abi=./tmp/abigen/erc1155 --bin=./tmp/bytecode/erc1155 --pkg erc1155 --out=abis/token/erc1155/erc1155.go
abigen --abi=./tmp/abigen/erc1155Supply --pkg erc1155Supply --out=abis/token/erc1155/extensions/erc1155Supply/erc1155Supply.go
abigen --abi=./tmp/abigen/erc1155MetadataURI --pkg erc1155Metadata --out=abis/token/erc1155/extensions/erc1155Metadata/erc1155Metadata.go

rm -rf tmp/abigen
rm -rf tmp/bytecode
