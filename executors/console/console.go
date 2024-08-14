package console

import (
	"fmt"
	"strings"

	"github.com/todoesverso/aos/command/models"
)

type ConsoleExecutor struct{}

func (ce ConsoleExecutor) Execute(cmd models.OSCommand) error {
	commandStr := fmt.Sprintf("%s %s", cmd.Executable, strings.Join(cmd.Arguments, " "))
	fmt.Println(commandStr)
	return nil
}
