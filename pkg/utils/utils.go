package utils

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/everesthack-incubator/vault-differ/pkg/commander"

	log "github.com/sirupsen/logrus"
)

// FileList is a function that takes in the pattern and returns an array of matched files over a channel
func FileList(globPattern string, c chan []string, wg *sync.WaitGroup, logLevel log.Level) {
	defer wg.Done()
	log.SetLevel(logLevel)
	files, err := filepath.Glob(globPattern)
	log.Debugf("matched files from provided pattern: %s", files)
	if err != nil {
		log.Error("error processing the pattern:", err)
	}
	c <- files
}

func GitDiffCommands(filename string, c chan string, wg *sync.WaitGroup, logLevel log.Level) {
	defer wg.Done()
	log.SetLevel(logLevel)
	args := []string{"diff"}
	out, err := commander.Commands("git", args, filename)
	if err != nil {
		log.Error("error diffing files, are the secrets encrypted?:", err)
	}

	if len(out) == 0 {
		c <- filename
		log.Debugf("restorable file: %s", filename)
	}

}

func GitRestoreCommands(filename string, c chan bool, wg *sync.WaitGroup, logLevel log.Level, m *sync.RWMutex) {
	defer wg.Done()
	log.SetLevel(logLevel)
	args := []string{"restore", filename}
	cmd := exec.Command("git", args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	m.Lock()
	err := cmd.Run()
	m.Unlock()
	if err != nil {
		log.Debugf("failed to restore file: %s, %s", err, args)
	}

	log.Debugf("restore command output/error/stdout/stderr: %s , %s, %s, %s", cmd, err, stdout.Bytes(), stderr.Bytes())
	c <- true
}

func VaultToolCmd(file string, wg *sync.WaitGroup, cmd string, args []string, c chan bool, m *sync.RWMutex) {
	defer wg.Done()
	log.Debugln("commands/arguments/files", cmd, args, file)
	m.Lock()
	commander.Commands(cmd, args, file)
	m.Unlock()
	c <- true
}
