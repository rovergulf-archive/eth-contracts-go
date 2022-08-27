package signer

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/rovergulf/eth-contracts-go/pkg/ethutils"
	"github.com/spf13/viper"
)

// DefaultSigner prepares and returns specified private eth key as keystore.Key
func DefaultSigner() (*keystore.Key, error) {
	pk := viper.GetString("private_key")
	if len(pk) == 0 {
		return nil, fmt.Errorf("private key not specified")
	}

	k, err := ethutils.PrivateKeyStringToKey(pk)
	if err != nil {
		return nil, err
	}

	return k, nil
}
