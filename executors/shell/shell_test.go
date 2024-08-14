package shell

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/todoesverso/aos/command/models"
)

func TestShellExecutor_Execute_Success(t *testing.T) {
	executor := ShellExecutor{}

	cmd := models.OSCommand{
		Executable: "echo",
		Arguments:  []string{"Hello, World!"},
	}

	if err := executor.Execute(cmd); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestShellExecutor_Execute_CommandFailure(t *testing.T) {
	executor := ShellExecutor{}

	cmd := models.OSCommand{
		Executable: "invalid_command",
		Arguments:  []string{},
	}

	if err := executor.Execute(cmd); err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestShellExecutor_Execute_OutputCapture(t *testing.T) {
	executor := ShellExecutor{}

	cmd := models.OSCommand{
		Executable: "echo",
		Arguments:  []string{"capture this"},
	}

	// Capture stdout
	r, w, _ := os.Pipe()
	originalStdout := os.Stdout
	os.Stdout = w

	var buf bytes.Buffer
	done := make(chan struct{})

	go func() {
		io.Copy(&buf, r)
		close(done)
	}()

	if err := executor.Execute(cmd); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Close writer and restore stdout
	w.Close()
	os.Stdout = originalStdout
	<-done

	expected := "capture this\n"
	if buf.String() != expected {
		t.Errorf("expected %q, got %q", expected, buf.String())
	}
}
