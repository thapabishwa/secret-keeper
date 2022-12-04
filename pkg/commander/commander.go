package commander

import (
	"os/exec"

	log "github.com/sirupsen/logrus"
)

type Runner interface {
	CombinedOutput() ([]byte, error)
}

type Commander struct {
	*exec.Cmd
}

func NewCommander(command string, args []string, filename interface{}) Runner {
	// check if filename is a string or a slice of strings
	var newargs []string
	switch f := filename.(type) {
	case string:
		newargs = append(args, f)
	case []string:
		newargs = append(args, f...)
	}
	cmd := exec.Command(command, newargs...)
	return &Commander{cmd}
}

var ExecCommander = NewCommander

func Command(command string, args []string, filename interface{}) ([]byte, error) {
	out, err := ExecCommander(command, args, filename).CombinedOutput()
	if err != nil {
		log.Debugf("error running commands: %s, %s", err, string(out))
	}
	return out, err
}

func GitDiff(filename string) ([]byte, error) {
	out, err := ExecCommander("git", []string{"diff"}, filename).CombinedOutput()
	if err != nil {
		log.Debugf("error running commands: %s, %s", err, string(out))
	}
	return out, err
}

func GitRestore(files []string) ([]byte, error) {
	out, err := ExecCommander("git", []string{"restore"}, files).CombinedOutput()
	if err != nil {
		log.Debugf("error running commands: %s, %s", err, string(out))
	}
	return out, err
}
