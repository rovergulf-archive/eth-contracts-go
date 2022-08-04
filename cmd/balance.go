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
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/rovergulf/eth-contracts-go/abis/token/erc20"
	"github.com/spf13/cobra"
	"math/big"
)

func init() {
	rootCmd.AddCommand(balanceCmd)

	balanceCmd.Flags().StringP("address", "a", "", "Specify address")
	balanceCmd.Flags().StringP("token", "t", "", "ERC20/ERC777 token address")
}

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Check address balance",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		addr, err := cmd.Flags().GetString("address")
		if err != nil {
			return err
		}

		if !common.IsHexAddress(addr) {
			return errors.New("invalid address")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		strAddress, _ := cmd.Flags().GetString("address")
		token, _ := cmd.Flags().GetString("token")

		addr := common.HexToAddress(strAddress)

		client, err := ethclient.DialContext(ctx, providerUrl)
		if err != nil {
			logger.Errorw("Unable to connect provider", "err", err)
			return err
		}

		if len(token) > 0 {
			if !common.IsHexAddress(token) {
				return errors.New("invalid token address")
			}

			tokenInstance, err := erc20.NewErc20(common.HexToAddress(token), client)
			if err != nil {
				logger.Errorw("Unable to get ERC20 instance", "err", err)
				return err
			}

			balance, err := tokenInstance.BalanceOf(nil, addr)
			if err != nil {
				logger.Errorw("Unable to retrieve address token balance", "err", err)
				return err
			}

			logger.Infow("Balance",
				"address", addr,
				"token", token,
				"wei_balance", balance,
			)
		} else {
			blockNumber, err := client.BlockNumber(ctx)
			if err != nil {
				return err
			}

			bn := new(big.Int)
			bn.SetUint64(blockNumber)
			balance, err := client.BalanceAt(ctx, addr, bn)
			if err != nil {
				logger.Errorw("Unable to retrieve address balance", "err", err)
				return err
			}

			logger.Infow("Balance",
				"address", addr,
				"wei_balance", balance,
				"eth_balance", new(big.Int).Div(balance, big.NewInt(params.Ether)),
			)
		}
		return nil
	},
}
