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
	vaultInstance.MatchFiles()
	if len(vaultInstance.encryptArgs) == 0 || vaultInstance.vaultTool == "" {
		log.Fatal("vault tool not defined properly")
	} else {
		vaultInstance.Encrypt()
		vaultInstance.Differ()
		vaultInstance.Clean()
	}

}
