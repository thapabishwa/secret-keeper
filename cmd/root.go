package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "vault-differ",
		Short: "A tool to manage vault secrets",
		Long:  `vault-differ is a CLI tool that complements tools like ansible-vault, sops and more. It helps in managing and storing secrets in git-based repositories.`,
	}

	vaultInstance = NewVaultDiffer()

	config = NewConfig()
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ansible-vault-differ/config.yaml)")

}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

func initConfig() {

	log.SetOutput(os.Stdout)

	// set config type to yaml
	viper.SetConfigType("yaml")
	// search for config in /etc/
	viper.AddConfigPath("/etc/vault-differ/")
	// search for config in home dir
	viper.AddConfigPath("$HOME/.vault-differ/")
	// search for config in current dir
	viper.AddConfigPath(".")
	// config file name
	viper.SetConfigName("config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found
			log.Fatal("config file not found")
		} else {
			// Config file was found but another error was produced
			log.Fatal("some other error occured", ok)
		}
	}

	err := viper.Unmarshal(config)

	if err != nil {
		log.Fatal("cannot unmarshal config file")
	}

	vaultInstance.InitConfig(*config)

}
