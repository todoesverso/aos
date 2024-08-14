package helper

import (
	"fmt"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/todoesverso/aos/command/models"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

type HelperExecutor struct{ YamlInput inputmodels.YamlInput }

func (he HelperExecutor) Execute(cmd models.OSCommand) error {
	if he.YamlInput.Description == "" {
		fmt.Println("This AoS has no description text, sorry.")
		return nil
	}

	lines := strings.Split(he.YamlInput.Description, "\n")
	var wrappedText string
	for _, line := range lines {
		wrappedText += pterm.DefaultParagraph.WithMaxWidth(60).Sprintln(line)
	}
	title := pterm.LightRed(os.Args[1])
	pterm.DefaultBox.WithTitle(title).Println(wrappedText)
	return nil
}
