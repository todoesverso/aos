package shell

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/todoesverso/aos/command/models"
)

type ShellExecutor struct{}

func (se ShellExecutor) Execute(oscmd models.OSCommand) error {
	cmd := exec.Command(oscmd.Executable, oscmd.Arguments...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("failed to obtain stdout: %w", err)
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("failed to obtain stderr: %w", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start command: %w", err)
	}

	errChan := make(chan error, 2)

	go func() {
		if _, err := io.Copy(os.Stdout, stdout); err != nil {
			errChan <- fmt.Errorf("failed to copy stdout: %w", err)
		}
		errChan <- nil
	}()

	go func() {
		if _, err := io.Copy(os.Stderr, stderr); err != nil {
			errChan <- fmt.Errorf("failed to copy stderr: %w", err)
		}
		errChan <- nil
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("command execution failed: %w", err)
	}

	return nil
}
