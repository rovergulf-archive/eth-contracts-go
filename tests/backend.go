package tests

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"math/big"
	"os"
)

var (
	TestProviderUrl = os.Getenv("TEST_ETH_PROVIDER_URL")
	defaultAlloc    = new(big.Int)
)

func NewFakeBackend() *backends.SimulatedBackend {
	defaultAlloc.SetString("1000000000000000000000", 10)
	return backends.NewSimulatedBackend(core.GenesisAlloc{
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
	}, 4712388)
}
