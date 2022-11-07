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
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/rovergulf/eth-contracts-go/watcher"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(watchCmd)

	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")
	// watchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Subscribe to specified contract log events",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()

		var addrs []common.Address
		var topics [][]common.Hash

		w, err := watcher.New(ctx, providerUrl)
		if err != nil {
			return err
		}

		return w.WatchEvents(ctx, addrs, topics, func(ctx context.Context, eventLog types.Log) {
			fmt.Println(eventLog)
		})
	},
}
