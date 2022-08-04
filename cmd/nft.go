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
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(nftCmd)

	// nftCmd.PersistentFlags().String("foo", "", "A help for foo")
	// nftCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	nftCmd.AddCommand(nftAssetsCmd)
}

// nftCmd represents the nft command
var nftCmd = &cobra.Command{
	Use:          "nft",
	Short:        "NFT API extension",
	Long:         ``,
	SilenceUsage: true,
}

// nftAssetsCmd represents the assets command
var nftAssetsCmd = &cobra.Command{
	Use:   "assets",
	Short: "Discover specified address assets",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Info("owner-assets called")
		return nil
	},
}
