package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Removes unchanged secrets from git repositories",
	Long:  "This command compares the diff between current change and the HEAD and restores the original file if the secrets were not actually changed",
	Run:   cleanCmdRun,
}

var cleanCmdRun = func(cmd *cobra.Command, args []string) {
	matchedFiles := vaultInstance.MatchFiles()
	diffedFiles := vaultInstance.Differ(matchedFiles)
	cleanFiles := vaultInstance.Clean(diffedFiles)
	for file := range cleanFiles {
		log.Debug("cleaned file: ", file)
	}
}
