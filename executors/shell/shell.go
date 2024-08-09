package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"

  command "github.com/todoesverso/aos/command/models"
)

type ShellExecutor struct{}

func (se ShellExecutor) Execute(oscmd command.OSCommand) error {
	cmd := exec.Command(oscmd.Executable, oscmd.Arguments...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("Error obtaining stdout: %v", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("Error obtaining stderr: %v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("Error starting command: %v", err)
	}

	go func() error {
		if _, err := io.Copy(os.Stdout, stdout); err != nil {
			return fmt.Errorf("Error copying stdout: %v", err)
		}
		return nil
	}()

	go func() error {
		if _, err := io.Copy(os.Stderr, stderr); err != nil {
			return fmt.Errorf("Error copying stderr: %v", err)
		}
		return nil
	}()

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("Error waiting for command: %v", err)
	}

	return nil
}
