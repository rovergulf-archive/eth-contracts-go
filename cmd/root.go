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
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rovergulf/eth-contracts-go/pkg/logutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"os"
)

var (
	cfgFile     string
	providerUrl string
	logger      *zap.SugaredLogger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eth-contracts-go",
	Short: "CLI application for Ethereum contracts interaction",
	Long:  ``,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if logger != nil {
			logger.Error(err)
		} else {
			log.Println(err)
		}
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eth-contracts-go.yaml)")
	rootCmd.PersistentFlags().StringVar(&providerUrl, "provider-url", "", "Ethereum provider url. Uses ETH_PROVIDER_URL env as default")
	rootCmd.PersistentFlags().String("private-key", "", "Ethereum private signer key. Uses ETHEREUM_PRIVATE_KEY env as default")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("env", "e", "dev", "Environment to use. Not actually does anything")
	rootCmd.Flags().Bool("log-json", false, "Enable json logger output")
	rootCmd.Flags().Bool("log-stacktrace", false, "Enable logger error stacktrace output")

	viper.BindPFlag("provider_url", rootCmd.PersistentFlags().Lookup("provider-url"))
	viper.BindPFlag("env", rootCmd.Flags().Lookup("env"))
	viper.BindPFlag("log_json", rootCmd.Flags().Lookup("log-json"))
	viper.BindPFlag("log_stacktrace", rootCmd.Flags().Lookup("log-stacktrace"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetDefault("jaeger_trace", os.Getenv("JAEGER_TRACE_COLLECTOR"))
	viper.SetDefault("provider_url", os.Getenv("ETH_PROVIDER_URL"))
	viper.SetDefault("eth_private_key", os.Getenv("ETHEREUM_PRIVATE_KEY"))

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".eth-contracts-go" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(os.Getenv("ETH_CONTRACTS_GO_HOME"))
		viper.SetConfigName(".eth-contracts-go")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
		cfgFile = viper.ConfigFileUsed()
	}

	if err := initLogger(); err != nil {
		log.Fatal(err)
	}
}

func initLogger() error {
	var err error
	logger, err = logutils.NewLogger()
	return err
}

func addAndBindPrivateKeyFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("private-key", "pk", "", "Specify private key path or string value")
	viper.BindPFlag("eth_private_key", cmd.Flags().Lookup("private_key"))
	//cmd.Flags().Bool("password", false, "Type password to unlock private key file")
}

func addAddressFlag(cmd *cobra.Command) {
	cmd.Flags().StringP("address", "a", "", "Specify address")
}

func checkIfAddressProvided(cmd *cobra.Command, args []string) error {
	addr, err := cmd.Flags().GetString("address")
	if err != nil {
		return err
	}

	if !common.IsHexAddress(addr) {
		return errors.New("invalid address")
	}

	return nil
}
