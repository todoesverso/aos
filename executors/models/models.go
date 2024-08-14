package models

import (
	command "github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/executors/console"
	"github.com/todoesverso/aos/executors/explainer"
	"github.com/todoesverso/aos/executors/helper"
	"github.com/todoesverso/aos/executors/shell"
	"github.com/todoesverso/aos/executors/usage"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

type EnvOption byte

const (
	Charh EnvOption = 'h'
	CharH EnvOption = 'H'
	CharR EnvOption = 'R'
	CharX EnvOption = 'X'
	CharE EnvOption = 'E'
	Chard EnvOption = 'd' // DEFAULT EXECUTOR
)

func InitExecutors(YamlInput inputmodels.YamlInput) {
	RegisterExecutor(CharR, console.ConsoleExecutor{})
	RegisterExecutor(CharE, explainer.ExplainerExecutor{YamlInput: YamlInput})
	RegisterExecutor(Charh, usage.UsageExecutor{})
	RegisterExecutor(CharH, helper.HelperExecutor{YamlInput: YamlInput})
	RegisterExecutor(CharX, shell.ShellExecutor{})
	RegisterExecutor(Chard, shell.ShellExecutor{})
}

type Executor interface {
	Execute(cmd command.OSCommand) error
}

type CommandBuilder interface {
	Build(input inputmodels.YamlInput) command.OSCommand
}

type OutputProcessor interface {
	Executor
}

// Global executor registry
var ExecutorRegistry = make(map[EnvOption]Executor)

// RegisterExecutor allows registration of a new executor with a specific char
func RegisterExecutor(char EnvOption, executor Executor) {
	ExecutorRegistry[char] = executor
}
