package cmd

import (
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
	vaultInstance.MatchFiles()
	if len(vaultInstance.decryptArgs) == 0 || vaultInstance.vaultTool == "" {
		log.Fatalf("vault tools not defined properly")
	} else {
		vaultInstance.Decrypt()
	}
}
