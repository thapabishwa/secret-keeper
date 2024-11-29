package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the repo with the provided config file",
	Long:  "This command parses the config file and sets up the repo for secret-keeper",
	Run:   initCmdRun,
}

var initCmdRun = func(cmd *cobra.Command, args []string) {
	err := vaultInstance.BuildGitAttributes()
	if err != nil {
		log.Fatal(err)
	}

	err = vaultInstance.BuildGitConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = vaultInstance.AddPreCommitHook()
	if err != nil {
		log.Fatal(err)
	}
}
