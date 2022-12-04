package helpers

import (
	"bytes"
	"os/exec"
	"path/filepath"

	"github.com/everesthack-incubator/vault-differ/pkg/commander"

	log "github.com/sirupsen/logrus"
)

// FileList is a function that takes in the pattern and returns an array of matched files over a channel
func FileList(globPattern string, logLevel log.Level) []string {
	files, err := filepath.Glob(globPattern)
	log.Debugf("matched files from provided pattern: %s", files)
	if err != nil {
		log.Error("error processing the pattern:", err)
	}
	return files
}

func GitDiffCommands(filename string, logLevel log.Level) ([]byte, error) {
	log.SetLevel(logLevel)
	args := []string{"diff"}
	out, err := commander.Commands("git", args, filename)
	if err != nil {
		log.Error("error diffing files, are the secrets encrypted?:", err)
	}
	//log.Error("git diff output:", string(out), err)

	if len(out) == 0 {
		log.Debugf("restorable file: %s", filename)
	}
	return out, err
}

// Pass the list of files to be cleaned instead of each file. This fixes the issue where the file is not cleaned because of git lock.
func GitRestoreCommands(filename []string, logLevel log.Level) {
	log.SetLevel(logLevel)
	args := []string{"restore"}
	args = append(args, filename...)
	cmd := exec.Command("git", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Debugf("failed to restore file: %s, %s", err, args)
	}

	err = cmd.Wait()
	if err != nil {
		log.Debugf("failed to restore file: %s, %s", err, args)
	}

	log.Debugf("restore command output/error/stdout/stderr: %s , %s, %s, %s", cmd, err, stdout.Bytes(), stderr.Bytes())
}

func VaultToolCmd(cmd string, args []string, file string) {
	commander.Commands(cmd, args, file)
}
