package cmd

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(decryptCmd)
}

var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Runs the decrypt command provided in the config file",
	Long:  "This command compares the diff between current change and the HEAD and restores the original file if the secrets were not actually changed",
	Run:   decryptCmdRun,
}

var decryptCmdRun = func(cmd *cobra.Command, args []string) {

	matchedFiles := vaultInstance.MatchFiles()
	if len(vaultInstance.GetEncryptArgs()) == 0 || vaultInstance.GetVaultCommand() == "" {
		log.Fatalf("vault tools not defined properly")
		os.Exit(1)
	}
	decryptedFiles := vaultInstance.Decrypt(matchedFiles)

	for file := range decryptedFiles {
		log.Debugf("decrypted file:", file)
	}

}
