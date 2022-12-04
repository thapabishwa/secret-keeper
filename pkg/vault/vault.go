package vault

import (
	"sync"

	"github.com/everesthack-incubator/vault-differ/pkg/commander"
	"github.com/everesthack-incubator/vault-differ/pkg/config"
	"github.com/everesthack-incubator/vault-differ/pkg/helpers"

	log "github.com/sirupsen/logrus"
)

// VaultDiffer represents the config
type VaultDiffer struct {
	secrets     []string
	logLevel    log.Level
	vaultTool   string
	encryptArgs []string
	decryptArgs []string
}

// NewVaultDiffer returns an empty instance of VaultDiffer
func NewVaultDiffer() *VaultDiffer {
	return &VaultDiffer{}
}

func (a *VaultDiffer) GetEncryptArgs() []string {
	return a.encryptArgs
}

func (a *VaultDiffer) GetDecryptArgs() []string {
	return a.decryptArgs
}

func (a *VaultDiffer) GetVaultCommand() string {
	return a.vaultTool
}

// InitConfig Reads and Updates all config
func (a *VaultDiffer) InitConfig(config config.Config) {
	a.secrets = config.Secrets
	a.vaultTool = config.VaultTool
	a.encryptArgs = config.EncryptArgs
	a.decryptArgs = config.DecryptArgs
	a.logLevel = log.InfoLevel
	if config.Debug {
		a.logLevel = log.DebugLevel
	}
}

// MatchFiles populates list of files that match the pattern provided in the config
func (a *VaultDiffer) MatchFiles() <-chan string {
	processedFiles := make(chan string)
	go func() {
		for _, pattern := range a.secrets {
			files, err := helpers.FileList(pattern, a.logLevel)
			if err != nil {
				log.Error(err)
			}
			for _, file := range files {
				processedFiles <- file
			}
		}
		close(processedFiles)
	}()
	return processedFiles

}

func (a *VaultDiffer) Clean(files <-chan string) <-chan string {
	processedFiles := make(chan string)
	go func() {
		restorableFiles := []string{}
		for file := range files {
			restorableFiles = append(restorableFiles, file)
			processedFiles <- file
		}
		if len(restorableFiles) > 0 {
			gitRestore, err := commander.GitRestore(restorableFiles)
			if err != nil {
				log.Error(err, gitRestore)
			}
		}
		close(processedFiles)
	}()
	return processedFiles
}

func (a *VaultDiffer) Differ(files <-chan string) <-chan string {
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
						log.Debugf("error diffing file: %s, status code %s, %s", file, err.Error(), string(out))
					} else {
						log.Errorf("error diffing file: %s", file)
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
func (a *VaultDiffer) Encrypt(files <-chan string) <-chan string {
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
						log.Debugf("error decrypting file: %s, status code %s, %s", file, err.Error(), string(out))
					} else {
						log.Errorf("error decrypting file: %s", file)
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
func (a *VaultDiffer) Decrypt(files <-chan string) <-chan string {
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
						log.Debugf("error decrypting file: %s, status code %s, %s", file, err.Error(), string(out))
					} else {
						log.Errorf("error decrypting file: %s", file)
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
