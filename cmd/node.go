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
	"github.com/rovergulf/eth-contracts-go/tests"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	rootCmd.AddCommand(nodeCmd)
}

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Start local eth node",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		//ctx := context.Background()
		quitChan := make(chan os.Signal, 1)
		signal.Notify(quitChan, os.Interrupt, os.Kill, syscall.SIGTERM)

		gethNode, err := tests.NewFakeGethNode()
		if err != nil {
			return err
		}

		if err := gethNode.Eth.Start(); err != nil {
			return err
		}

		<-quitChan
		return gethNode.Eth.Stop()
	},
}
