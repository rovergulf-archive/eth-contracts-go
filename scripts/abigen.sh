#!/usr/bin/env bash

set -e

source ./scripts/util.sh

mkdir -p tmp/abigen
jsonnet hack/jsonnet/abi.jsonnet -m tmp/abigen

abigen --abi=./tmp/abigen/erc20 --pkg erc20 --out=abis/erc20/erc20.go
abigen --abi=./tmp/abigen/erc165 --pkg erc165 --out=abis/erc165/erc165.go
abigen --abi=./tmp/abigen/erc721 --pkg erc721 --out=abis/erc721/erc721.go
abigen --abi=./tmp/abigen/erc721Enumerable --pkg erc721Enumerable --out=abis/erc721/extensions/erc721Enumerable/erc721Enumerable.go
abigen --abi=./tmp/abigen/erc721Metadata --pkg erc721Metadata --out=abis/erc721/extensions/erc721Metadata/erc721Metadata.go
abigen --abi=./tmp/abigen/erc777 --pkg erc777 --out=abis/erc777/erc777.go
abigen --abi=./tmp/abigen/erc1155 --pkg erc1155 --out=abis/erc1155/erc1155.go
abigen --abi=./tmp/abigen/erc1155Supply --pkg erc1155Supply --out=abis/erc1155/extensions/erc1155Supply/erc1155Supply.go
abigen --abi=./tmp/abigen/erc1155MetadataURI --pkg erc1155Metadata --out=abis/erc1155/extensions/erc1155Metadata/erc1155Metadata.go

rm -rf tmp/abigen
