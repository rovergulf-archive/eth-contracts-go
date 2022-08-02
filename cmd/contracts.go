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
	"github.com/rovergulf/eth-contracts-go/pkg/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(contractsCmd)
	//contractsCmd.PersistentFlags().StringP("interface-id", "id", "", "ERC Interface ID name (erc165, erc20, etc...)")

	contractsCmd.AddCommand(contractsCheckInterfaceIdCmd)
	contractsCmd.AddCommand(contractsCallCmd)
	contractsCmd.AddCommand(contractsWatchCmd)
}

// contractsCmd represents the contracts command
var contractsCmd = &cobra.Command{
	Use:          "contracts",
	Short:        "Interact with Ethereum contracts",
	Long:         ``,
	SilenceUsage: true,
	//Run: func(cmd *cobra.Command, args []string) {
	//	logger.Info("contracts called")
	//},
}

// contractsInterfaceIdCmd represents the contracts command
var contractsCheckInterfaceIdCmd = &cobra.Command{
	Use:   "check-interface-id",
	Short: "Check contract interface id",
	Long:  `Checks contract interface id using ERC165`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewWithLogger(viper.GetString("provider_url"), logger)
		if err != nil {
			return err
		}
		defer c.Close()

		logger.Info("check interface id called")

		return nil
	},
}

// contractsInterfaceIdCmd represents the contracts command
var contractsCallCmd = &cobra.Command{
	Use:   "call",
	Short: "Call contract methods",
	Long: `Calls specified contract based on provided (and supported by this package) interface id and address
It will return error if call arguments are required`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("call called")
		return nil
	},
}

// contractsInterfaceIdCmd represents the contracts command
var contractsWatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch contract events",
	Long:  `Watch specified contract topics`,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("watch called")
		return nil
	},
}
