package common

import (
	"fmt"
	"os"

	"github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/inputs/argpars"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

const DEFAULT_ARGS_NUMBER = 2 // executable and input file
const POSITIONAL_ARG_SENTINEL = "__AOS_PA__"

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
	for _, arg := range input.Arguments {
		if arg.Positional != nil {
			if arg.Positional.Order < 0 {
				return fmt.Errorf("Malformed positional argument in input file. Order must be bigger than 0.")
			}
			if arg.Positional.Order-1 >= positionalCount {
				return fmt.Errorf("Malformed positional argument in input file. Order is to big.")
			}
		}
	}

	return nil
}

func (ccb CommonCommandBuilder) Build(input inputmodels.YamlInput) (models.OSCommand, error) {
	err := validatePositionalArgs(input)
	if err != nil {
		return models.OSCommand{}, err
	}

	arguments, positionalCount := ccb.processArguments(input.Arguments)
	orderedPositionalArgs := ccb.orderPositionalArgs(input.Arguments, positionalCount)
	finalArguments := ccb.buildFinalArgs(arguments, orderedPositionalArgs)

	return models.OSCommand{
		Executable: input.Command.Exec,
		Arguments:  finalArguments,
	}, nil
}

func (ccb CommonCommandBuilder) processArguments(args []inputmodels.Argument) ([]string, int) {
	var arguments []string
	positionalCount := 0

	for _, arg := range args {
		if arg.Positional != nil {
			arguments = append(arguments, POSITIONAL_ARG_SENTINEL)
			positionalCount++
		} else {
			arguments = append(arguments, argpars.ParseArgument(arg)...)
		}
	}

	return arguments, positionalCount
}

func (ccb CommonCommandBuilder) orderPositionalArgs(args []inputmodels.Argument, positionalCount int) []string {
	orderedPositionalArgs := make([]string, positionalCount)
	positionalIndex := 0

	for _, arg := range args {
		if arg.Positional != nil {
			pos := os.Args[positionalIndex+DEFAULT_ARGS_NUMBER]
			order := arg.Positional.Order
			if order > 0 {
				orderedPositionalArgs[order-1] = pos
			} else {
				orderedPositionalArgs[positionalIndex] = pos
			}
			positionalIndex++
		}
	}

	return orderedPositionalArgs
}

func (ccb CommonCommandBuilder) buildFinalArgs(args []string, orderArgs []string) []string {
	arguments := make([]string, len(args))
	positionalIndex := 0
	for i, arg := range args {
		if arg == POSITIONAL_ARG_SENTINEL {
			arguments[i] = orderArgs[positionalIndex]
			positionalIndex++
		} else {
			arguments[i] = arg
		}
	}

	return arguments

}
