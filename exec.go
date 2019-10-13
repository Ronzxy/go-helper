package helper

import (
	"io/ioutil"
	"os/exec"
)

var Exec = NewExecHelper()

type ExecHelper struct{}

func NewExecHelper() *ExecHelper {
	return &ExecHelper{}
}

func (this *ExecHelper) Run(name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer stdout.Close()

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(stdout)
}
