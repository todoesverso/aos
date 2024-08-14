package console

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/todoesverso/aos/command/models"
)

// Helper function to capture fmt.Println output
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestConsoleExecutor_Execute(t *testing.T) {
	tests := []struct {
		name     string
		cmd      models.OSCommand
		expected string
	}{
		{
			name:     "No arguments",
			cmd:      models.OSCommand{Executable: "echo", Arguments: []string{}},
			expected: "echo\n",
		},
		{
			name:     "With single argument",
			cmd:      models.OSCommand{Executable: "echo", Arguments: []string{"hello"}},
			expected: "echo hello\n",
		},
		{
			name:     "With multiple arguments",
			cmd:      models.OSCommand{Executable: "ls", Arguments: []string{"-l", "/home/user"}},
			expected: "ls -l /home/user\n",
		},
	}

	ce := ConsoleExecutor{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := captureOutput(func() {
				err := ce.Execute(tt.cmd)
				if err != nil {
					t.Fatalf("Execute() error = %v", err)
				}
			})

			if strings.TrimSpace(output) != strings.TrimSpace(tt.expected) {
				t.Errorf("Execute() = %v, want %v", output, tt.expected)
			}
		})
	}
}
