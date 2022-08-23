/*
Copyright © 2022 Rovergulf Engineers Team <team@rovergulf.net>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rovergulf/eth-contracts-go/nft"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nftCmd)

	// nftCmd.PersistentFlags().String("foo", "", "A help for foo")
	// nftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	nftCmd.AddCommand(nftAssetsCmd)
	addAddressFlag(nftAssetsCmd)

	nftCmd.AddCommand(nftInfoCmd)
	addAddressFlag(nftInfoCmd)
}

// nftCmd represents the nft command
var nftCmd = &cobra.Command{
	Use:          "nft",
	Short:        "NFT API extension",
	Long:         ``,
	SilenceUsage: true,
}

// nftAssetsCmd represents the assets command
var nftInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Discover specified NFT address contract information",
	Long:  ``,
	Args:  checkIfAddressProvided,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		strAddress, _ := cmd.Flags().GetString("address")
		tokenAddr := common.HexToAddress(strAddress)

		client, err := ethclient.DialContext(ctx, providerUrl)
		if err != nil {
			return err
		}
		defer client.Close()

		info, err := nft.FindCollectionMetadata(ctx, client, tokenAddr)
		if err != nil {
			return err
		}

		logger.Infow("NFT Contract verified",
			"interface", info.InterfaceID,
			"name", info.Name,
			"symbol", info.Symbol,
		)
		return nil
	},
}

// nftAssetsCmd represents the assets command
var nftAssetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Discover assets owned by specified address",
	Long:  ``,
	Args:  checkIfAddressProvided,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		strAddress, _ := cmd.Flags().GetString("address")
		strOwner, err := cmd.Flags().GetString("owner")
		tokenAddr := common.HexToAddress(strAddress)
		tokensOwner := common.HexToAddress(strOwner)

		client, err := ethclient.DialContext(ctx, providerUrl)
		if err != nil {
			return err
		}
		defer client.Close()

		assets, err := nft.FindCollectionOwnerAssets(ctx, client, tokenAddr, tokensOwner)
		if err != nil {
			logger.Errorw("Unable to find owner assets", "err", err)
			return err
		}

		if assets.AssetsCount > 0 {
			logger.Infof("%s is an owner of %d assets", tokensOwner, assets.AssetsCount)
		} else {
			logger.Info("No assets found for specified address")
		}

		return nil
	},
}
