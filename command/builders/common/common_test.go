package common

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todoesverso/aos/command/models"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

func TestCommonCommandBuilder_Build(t *testing.T) {
	tests := []struct {
		name          string
		input         inputmodels.YamlInput
		args          []string
		expectedCmd   models.OSCommand
		expectedError error
	}{
		{
			name: "Valid input with positional and optional arguments",
			input: inputmodels.YamlInput{
				Command: inputmodels.Command{
					Exec: "echo",
				},
				Arguments: []inputmodels.Argument{
					{Positional: &inputmodels.PositionalArgument{}},
					{Option: "-n",
						Value: "value",
					},
				},
			},
			args: []string{"cmd", "input.yaml", "positionalArg"},
			expectedCmd: models.OSCommand{
				Executable: "echo",
				Arguments:  []string{"positionalArg", "-n", "value"},
			},
			expectedError: nil,
		},
		{
			name: "Invalid input with missing positional argument",
			input: inputmodels.YamlInput{
				Command: inputmodels.Command{
					Exec: "echo",
				},
				Arguments: []inputmodels.Argument{
					{Positional: &inputmodels.PositionalArgument{}},
					{Positional: &inputmodels.PositionalArgument{}},
				},
			},
			args:          []string{"cmd", "input.yaml", "positionalArg1"},
			expectedCmd:   models.OSCommand{},
			expectedError: fmt.Errorf("Malformed command, 2 positional arguments declared and 1 provided"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock os.Args for testing
			originalArgs := os.Args
			defer func() { os.Args = originalArgs }() // Restore original os.Args after test
			os.Args = tt.args

			builder := CommonCommandBuilder{}
			cmd, err := builder.Build(tt.input)

			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedCmd, cmd)
		})
	}
}
