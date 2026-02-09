// Package utils provides utility functions
package utils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ExecuteCommand runs a command with the given arguments and returns stdout, stderr, and error
// This is a convenience wrapper around exec.Command that captures both output streams
func ExecuteCommand(name string, args ...string) (string, string, error) {
	var stdout, stderr bytes.Buffer

	cmd := exec.Command(name, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), fmt.Errorf("command '%s' failed: %w", name, err)
	}

	return stdout.String(), stderr.String(), nil
}
