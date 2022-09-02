package tests

import (
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/eth/catalyst"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/miner"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	"time"
)

type GethNode struct {
	Eth *eth.Ethereum
	api catalyst.ConsensusAPI
}

func NewFakeGethNode() (*GethNode, error) {
	// import genesis
	gen := &core.Genesis{
		Config: params.AllEthashProtocolChanges,
		Alloc: core.GenesisAlloc{
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
		},
		GasLimit: 4712388,
	}

	n, err := node.New(&node.Config{HTTPPort: 8545, AuthPort: 8445})
	if err != nil {
		panic(err)
	}

	ethcfg := &ethconfig.Config{
		Genesis:        gen,
		TrieTimeout:    time.Minute,
		TrieDirtyCache: 256,
		TrieCleanCache: 256,
		Miner:          miner.Config{GasCeil: 30_000_000},
	}
	ethservice, err := eth.New(n, ethcfg)
	if err != nil {
		return nil, err
	}

	return &GethNode{
		Eth: ethservice,
		api: *catalyst.NewConsensusAPI(ethservice),
	}, nil
}
