package cmd

import (
	"sync"

	log "github.com/sirupsen/logrus"
)

//VaultDiffer represents the config
type VaultDiffer struct {
	secrets          []string
	logLevel         log.Level
	matchedFiles     []string
	restoreableFiles []string
	vaultTool        string
	encryptArgs      []string
	decryptArgs      []string
}

// NewVaultDiffer returns an empty instance of VaultDiffer
func NewVaultDiffer() *VaultDiffer {
	return &VaultDiffer{}
}

// InitConfig Reads and Updates all config
func (a *VaultDiffer) InitConfig(config Config) {
	a.secrets = config.Secrets
	a.vaultTool = config.VaultTool
	a.encryptArgs = config.EncryptArgs
	a.decryptArgs = config.DecryptArgs
	if config.Debug {
		a.logLevel = log.DebugLevel
	}

}

// MatchFiles populates list of files that match the pattern provided in the config
func (a *VaultDiffer) MatchFiles() {
	c := make(chan []string, 100000)
	var wg sync.WaitGroup
	for i := 0; i < len(a.secrets); i++ {
		wg.Add(1)
		go FileList(a.secrets[i], c, &wg, a.logLevel)
	}
	wg.Wait()
	close(c)
	for elem := range c {
		a.matchedFiles = append(a.matchedFiles, elem...)
	}
}

// Differ Runs compares the files
func (a *VaultDiffer) Differ() {
	c := make(chan string, 100000)
	var wg sync.WaitGroup
	for _, elem := range a.matchedFiles {
		wg.Add(1)
		go gitDiffCommands(elem, c, &wg, a.logLevel)

	}
	wg.Wait()
	close(c)

	for elem := range c {
		a.restoreableFiles = append(a.restoreableFiles, elem)
	}

}

// Clean restores the list of restorables
func (a *VaultDiffer) Clean() {
	c := make(chan bool, 100000)
	var m sync.RWMutex
	var wg sync.WaitGroup
	for _, elem := range a.restoreableFiles {
		wg.Add(1)
		go gitRestoreCommands(elem, c, &wg, a.logLevel, &m)
	}
	wg.Wait()
	close(c)
}

// Encrypt all files
func (a *VaultDiffer) Encrypt() {
	c := make(chan bool, 100000)
	var wg sync.WaitGroup
	var m sync.RWMutex
	for _, elem := range a.matchedFiles {
		wg.Add(1)
		go vaultToolCmd(elem, &wg, a.vaultTool, a.encryptArgs, c, &m)
	}
	wg.Wait()
	close(c)
}

// Decrypt all files
func (a *VaultDiffer) Decrypt() {
	c := make(chan bool, 100000)
	var wg sync.WaitGroup
	var m sync.RWMutex
	for _, elem := range a.matchedFiles {
		wg.Add(1)
		go vaultToolCmd(elem, &wg, a.vaultTool, a.decryptArgs, c, &m)
	}
	wg.Wait()
	close(c)
}
