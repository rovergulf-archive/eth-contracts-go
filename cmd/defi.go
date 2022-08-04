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

// defiCmd represents the defi command
var defiCmd = &cobra.Command{
	Use:   "defi",
	Short: "DeFi API extensions",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("defi called")
	},
}

func init() {
	rootCmd.AddCommand(defiCmd)

	// defiCmd.PersistentFlags().String("foo", "", "A help for foo")
	// defiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
