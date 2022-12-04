package commander

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// Commands runs any arbitrary command
func Commands(command string, args []string, filename string) ([]byte, error) {
	newargs := append(args, filename)
	out, err := exec.Command(command, newargs...).CombinedOutput()
	log.Debugf("command: %s %s", command, args)
	log.Debugf("command output: %s", out)
	if err != nil {
		log.Error("error running commands:", err, string(out), command, args, filename)
	}
	return out, nil
}
