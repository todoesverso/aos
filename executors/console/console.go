package console

import (
	"fmt"
	"strings"

	"github.com/todoesverso/aos/command/models"
)

type ConsoleExecutor struct{}

func (y ConsoleExecutor) Execute(cmd models.OSCommand) error {
	fmt.Printf("%s %s\n", cmd.Executable, strings.Join(cmd.Arguments, " "))
	return nil
}
