/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
	"github.com/rovergulf/eth-contracts-go/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(providerCmd)
}

// providerCmd represents the provider command
var providerCmd = &cobra.Command{
	Use:   "provider-info",
	Short: "Get provider Chain info",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		c, err := client.NewWithLogger(viper.GetString("provider_url"), logger)
		if err != nil {
			return err
		}

		chainId, err := c.ChainID(ctx)
		if err != nil {
			return err
		}

		blockNumber, err := c.BlockNumber(ctx)
		if err != nil {
			return err
		}

		logger.Infow("Provider info",
			"chain_id", chainId.Uint64(),
			"block_number", blockNumber,
		)
		return nil
	},
}
