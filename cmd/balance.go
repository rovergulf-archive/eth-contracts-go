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
	"github.com/rovergulf/eth-contracts-go/defi"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"math/big"
)

func init() {
	rootCmd.AddCommand(balanceCmd)

	addAddressFlag(balanceCmd)
	balanceCmd.Flags().StringP("token", "t", "", "ERC20/ERC777 token address")
}

// balanceCmd represents the balance command
var balanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "Check address balance",
	Long:  ``,
	Args:  checkIfAddressProvided,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		ctx, span := otel.GetTracerProvider().Tracer("cli").Start(ctx, "balance")
		defer span.End()

		strAddress, _ := cmd.Flags().GetString("address")
		token, _ := cmd.Flags().GetString("token")

		addr := common.HexToAddress(strAddress)

		client, err := ethclient.DialContext(ctx, providerUrl)
		if err != nil {
			logger.Errorw("Unable to connect provider", "err", err)
			return err
		}
		defer client.Close()

		if len(token) > 0 {
			if !common.IsHexAddress(token) {
				return errors.New("invalid token address")
			}

			balance, err := defi.GetTokenAddressBalance(ctx, client, common.HexToAddress(token), addr)
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
			balance, err := defi.GetAddressBalance(ctx, client, addr)
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
