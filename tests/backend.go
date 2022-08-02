package tests

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
)

var defaultAlloc = hexutil.MustDecodeBig("100e27")

type SimulatedBackend struct {
	*backends.SimulatedBackend
}

func (b *SimulatedBackend) Shutdown() {
	b.Close()
}

func NewFakeBackend() *SimulatedBackend {
	return &SimulatedBackend{
		SimulatedBackend: backends.NewSimulatedBackend(core.GenesisAlloc{
			Account0:  core.GenesisAccount{Balance: defaultAlloc},
			Account1:  core.GenesisAccount{Balance: defaultAlloc},
			Account2:  core.GenesisAccount{Balance: defaultAlloc},
			Account3:  core.GenesisAccount{Balance: defaultAlloc},
			Account4:  core.GenesisAccount{Balance: defaultAlloc},
			Account5:  core.GenesisAccount{Balance: defaultAlloc},
			Account6:  core.GenesisAccount{Balance: defaultAlloc},
			Account7:  core.GenesisAccount{Balance: defaultAlloc},
			Account8:  core.GenesisAccount{Balance: defaultAlloc},
			Account9:  core.GenesisAccount{Balance: defaultAlloc},
			Account10: core.GenesisAccount{Balance: defaultAlloc},
		}, 32000),
	}
}
