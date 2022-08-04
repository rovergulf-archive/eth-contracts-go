#!/usr/bin/env bash

set -e

go test ./pkg/ethutils -cover
go test ./defi -cover
