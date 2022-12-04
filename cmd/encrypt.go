package cmd

import (
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(encryptCmd)
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Runs the encrypt command provided in the config file",
	Long:  "This command compares the diff between current change and the HEAD and restores the original file if the secrets were not actually changed",
	Run:   encryptCmdRun,
}

var encryptCmdRun = func(cmd *cobra.Command, args []string) {
	matchedFiles := vaultInstance.MatchFiles()
	if len(vaultInstance.GetEncryptArgs()) == 0 || vaultInstance.GetVaultCommand() == "" {
		log.Fatal("vault tool not defined properly")
	}
	encryptedFiles := vaultInstance.Encrypt(matchedFiles)
	restorableFiles := vaultInstance.Differ(encryptedFiles)
	restoredFiles := vaultInstance.Clean(restorableFiles)

	for file := range restoredFiles {
		log.Debugf("encrypted file:", file)
	}

}
