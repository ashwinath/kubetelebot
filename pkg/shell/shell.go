package shell

import (
	"fmt"
	"os/exec"
)

func RunShell(binary string, args ...string) (*string, error) {
	cmd := exec.Command(binary, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		// if there was any error, print it here
		return nil, fmt.Errorf("could not run command: %v", err)
	}
	output := string(out)
	return &output, nil
}
