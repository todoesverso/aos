package usage

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todoesverso/aos/command/models"
)

func captureOutput(f func()) string {
	var buf bytes.Buffer
	originalStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	outputChan := make(chan string)
	go func() {
		_, _ = buf.ReadFrom(r)
		outputChan <- buf.String()
	}()

	f()
	w.Close()
	os.Stdout = originalStdout
	return <-outputChan
}

func TestPrintUsage(t *testing.T) {
	// Mock os.Args to simulate the program name
	originalArgs := os.Args
	os.Args = []string{"aos"}
	defer func() {
		os.Args = originalArgs
	}()

	expectedOutput := `AliasOnSteroids

Usage:
	aos <alias.yaml> [positional arguments of the alias]
Options:
	In order to keep CLI arguments as straightforward as possible,
	options are passed thru the AOS environment variables.

	AOS=h   aos <alias.yaml>	# Prints this usage
	AOS=H   aos <alias.yaml>	# Prints a helper short description of the alias
	AOS=E   aos <alias.yaml>	# Prints a helper long description of the alias
	AOS=R   aos <alias.yaml>	# Renders the command and prints it to stdout
	[AOS=X] aos <alias.yaml>	# Runs the command in a shell.
`

	output := captureOutput(PrintUsage)
	assert.Equal(t, expectedOutput, output, "The usage output should match the expected format")
}

func TestUsageExecutor_Execute(t *testing.T) {
	// Mock os.Args to simulate the program name
	originalArgs := os.Args
	os.Args = []string{"aos"}
	defer func() {
		os.Args = originalArgs
	}()

	expectedOutput := `AliasOnSteroids

Usage:
	aos <alias.yaml> [positional arguments of the alias]
Options:
	In order to keep CLI arguments as straightforward as possible,
	options are passed thru the AOS environment variables.

	AOS=h   aos <alias.yaml>	# Prints this usage
	AOS=H   aos <alias.yaml>	# Prints a helper short description of the alias
	AOS=E   aos <alias.yaml>	# Prints a helper long description of the alias
	AOS=R   aos <alias.yaml>	# Renders the command and prints it to stdout
	[AOS=X] aos <alias.yaml>	# Runs the command in a shell.
`

	output := captureOutput(func() {
		ue := UsageExecutor{}
		err := ue.Execute(models.OSCommand{})
		assert.NoError(t, err, "Execute should not return an error")
	})

	assert.Equal(t, expectedOutput, output, "The Execute output should match the expected usage format")
}
