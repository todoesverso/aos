package common

import (
	"fmt"
	"os"

	"github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/inputs/argpars"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

const DEFAULT_ARGS_NUMBER = 2 // executable and input file

type CommandBuilder interface {
	Build(input inputmodels.YamlInput, extraArgsCount int) models.OSCommand
}

type CommonCommandBuilder struct{}

func validatePositionalArgs(input inputmodels.YamlInput) error {

  extraArgsCount := len(os.Args) - DEFAULT_ARGS_NUMBER
	positionalCount := 0
	for _, arg := range input.Arguments {
		if arg.Positional != nil {
			positionalCount += 1
		}
	}
	if positionalCount > extraArgsCount {
		return fmt.Errorf("Malformed command, %d positional arguments declared and %d provided", positionalCount, extraArgsCount)
	}
	return nil
}
func (ccb CommonCommandBuilder) Build(input inputmodels.YamlInput) (models.OSCommand, error) {
	var arguments []string

	err := validatePositionalArgs(input)
	if err != nil {
		return models.OSCommand{}, err
	}

	positionalCount := 0
	for _, arg := range input.Arguments {
		if arg.Positional != nil {
			pos := os.Args[positionalCount + DEFAULT_ARGS_NUMBER]
			arguments = append(arguments, pos)
      positionalCount += 1
		} else {
			args := argpars.ParseArgument(arg)
			arguments = append(arguments, args...)
		}
	}
	return models.OSCommand{
		Executable: input.Command.Exec,
		Arguments:  arguments,
	}, nil
}
