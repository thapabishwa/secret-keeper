package secretkeeper

import (
	"sync"

	"github.com/everesthack-incubator/secret-keeper/pkg/commander"
	"github.com/everesthack-incubator/secret-keeper/pkg/config"
	"github.com/everesthack-incubator/secret-keeper/pkg/helpers"

	log "github.com/sirupsen/logrus"
)

// SecretKeeper represents the config
type SecretKeeper struct {
	filePatterns []string
	logLevel     log.Level
	vaultTool    string
	encryptArgs  []string
	decryptArgs  []string
}

// NewSecretKeeper returns an empty instance of VaultDiffer
func NewSecretKeeper() *SecretKeeper {
	return &SecretKeeper{}
}

func (a *SecretKeeper) GetEncryptArgs() []string {
	return a.encryptArgs
}

func (a *SecretKeeper) GetDecryptArgs() []string {
	return a.decryptArgs
}

func (a *SecretKeeper) GetVaultCommand() string {
	return a.vaultTool
}

// InitConfig Reads and Updates all config
func (a *SecretKeeper) InitConfig(config config.Config) {
	a.filePatterns = config.FilePatterns
	a.vaultTool = config.VaultTool
	a.encryptArgs = config.EncryptArgs
	a.decryptArgs = config.DecryptArgs
	a.logLevel = log.InfoLevel
	if config.Debug {
		a.logLevel = log.DebugLevel
	}
	log.SetLevel(a.logLevel)
}

// MatchFiles populates list of files that match the pattern provided in the config
func (a *SecretKeeper) MatchFiles() <-chan string {
	processedFiles := make(chan string)
	go func() {
		for _, pattern := range a.filePatterns {
			files, err := helpers.FileList(pattern)
			if a.logLevel == log.DebugLevel {
				log.Debugf("files matching pattern: %s, %v", pattern, files)
			}
			if err != nil {
				if a.logLevel == log.DebugLevel {
					log.Error("error getting file list", err, files, pattern, a.filePatterns)
				} else {
					log.Error("error getting file list", err)
				}
			}
			for _, file := range files {
				// Ignore config.secret-keeper.yaml file from being encrypted
				if file != "config.secret-keeper.yaml" {
					processedFiles <- file
				}
			}
		}
		close(processedFiles)
	}()
	return processedFiles
}

func (a *SecretKeeper) Clean(files <-chan string) <-chan string {
	processedFiles := make(chan string)
	go func() {
		restorableFiles := []string{}
		for file := range files {
			exists, err := commander.GitLog(file)
			if err != nil {
				if a.logLevel == log.DebugLevel {
					log.Errorf("error running git log on file: %s, status code %s, %s", file, err.Error(), string(exists))
				} else {
					log.Errorf("error cleaning file: %s", file)
				}
			}
			if len(exists) > 0 {
				restorableFiles = append(restorableFiles, file)
				processedFiles <- file
			}
		}

		if len(restorableFiles) > 0 {
			log.Infof("restoring file: %v to previous state as they were not changed", restorableFiles)
			output, err := commander.GitRestore(restorableFiles)
			if err != nil {
				if a.logLevel == log.DebugLevel {
					log.Errorf("error restoring files: %s, status code %s, %s", restorableFiles, err.Error(), string(output))
				} else {
					log.Errorf("error restoring files: %s", restorableFiles)
				}
			}
		}

		close(processedFiles)
	}()
	return processedFiles
}

func (a *SecretKeeper) Differ(files <-chan string) <-chan string {
	processedFiles := make(chan string)
	go func() {
		var wg sync.WaitGroup
		for file := range files {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				out, err := commander.GitDiff(file)
				if err != nil {
					if a.logLevel == log.DebugLevel {
						log.Errorf("error checking diff for file: %s, status code %s, %s", file, err.Error(), string(out))
					} else {
						log.Errorf("error checking diff for file: %s\n%s", file, string(out))
					}
				}
				if len(out) == 0 {
					processedFiles <- file
				}
			}(file)
		}
		wg.Wait()
		close(processedFiles)
	}()
	return processedFiles
}

// Encrypt all files
func (a *SecretKeeper) Encrypt(files <-chan string) <-chan string {
	processedFiles := make(chan string)
	go func() {
		var wg sync.WaitGroup
		for file := range files {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				out, err := commander.Command(a.vaultTool, a.encryptArgs, file)
				if err != nil {
					if a.logLevel == log.DebugLevel {
						log.Errorf("error encrypting file: %s, status code %s, %s", file, err.Error(), string(out))
					} else {
						log.Errorf("error encrypting file: %s \n%s", file, string(out))
					}
				} else {
					processedFiles <- file
				}
			}(file)
		}
		wg.Wait()
		close(processedFiles)
	}()
	return processedFiles
}

// Decrypt all files
func (a *SecretKeeper) Decrypt(files <-chan string) <-chan string {
	processedFiles := make(chan string)

	go func() {
		var wg sync.WaitGroup
		for file := range files {
			wg.Add(1)
			go func(file string) {
				defer wg.Done()
				out, err := commander.Command(a.vaultTool, a.decryptArgs, file)
				if err != nil {
					if a.logLevel == log.DebugLevel {
						log.Errorf("error decrypting file: %s, status code %s, %s", file, err.Error(), string(out))
					} else {
						log.Errorf("error decrypting file: %s\n%s", file, string(out))
					}
				} else {
					processedFiles <- file
				}
			}(file)
		}

		wg.Wait()
		close(processedFiles)
	}()
	return processedFiles
}
