package models

import (
	"github.com/todoesverso/aos/inputs/models"
	command "github.com/todoesverso/aos/command/models"
)

type Executor interface {
	execute(cmd command.OSCommand) error
}

type CommandBuilder interface {
	build(input models.YamlInput) command.OSCommand
}

type OutputProcessor interface {
	Executor
}

type CommonCommandBuilder struct{}

type ShellExecutor struct{}


