package cmd

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/thapabishwa/secret-keeper/pkg/config"
	"github.com/thapabishwa/secret-keeper/pkg/secretkeeper"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "secret-keeper",
		Short: "A tool to manage vault secrets",
		Long:  `secret-keeper is a CLI tool that complements tools like ansible-vault, sops and more. It helps in managing and storing secrets in git-based repositories.`,
	}

	vaultInstance  = secretkeeper.NewSecretKeeper()
	configurations = config.NewConfig()
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.secret-keeper/config.secret-keeper.yaml)")
}

func initConfig() {
	log.SetOutput(os.Stdout)

	// set config type to yaml
	viper.SetConfigType("yaml")
	// search for config in /etc/
	viper.AddConfigPath("/etc/secret-keeper/")
	// search for config in home dir

	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		log.Println("Not inside a Git repoisitory, falling back to default config path")
		viper.AddConfigPath("$HOME/.secret-keeper/")
		// search for config in current dir
		viper.AddConfigPath(".")
	} else {
		// Trim output and set the repository root as a config path
		repoRoot := strings.TrimSpace(string(output))
		log.Printf("Found Git repository root: %s", repoRoot)
		viper.AddConfigPath(repoRoot)
	}

	// config file name
	viper.SetConfigName("config.secret-keeper")
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

	err = viper.Unmarshal(configurations)
	if err != nil {
		log.Fatal("cannot unmarshal config file")
	}

	vaultInstance.InitConfig(*configurations)
}
