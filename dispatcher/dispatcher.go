package dispatcher

import (
	"fmt"
	"os"
  "strings"

	"github.com/todoesverso/aos/command/builders/common"
	"github.com/todoesverso/aos/executors/console"
	"github.com/todoesverso/aos/executors/explainer"
	command "github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/executors/shell"
	inputmodels "github.com/todoesverso/aos/inputs/models"
	"github.com/todoesverso/aos/usage"

	"github.com/pterm/pterm"
)

func printAOSDescription(YamlInput inputmodels.YamlInput) {
	if YamlInput.Description == "" {
		fmt.Println("This AoS has no description text, sorry.")
	} else {
		lines := strings.Split(YamlInput.Description, "\n")
		var wrappedText string
		for _, line := range lines {
			wrappedText += pterm.DefaultParagraph.WithMaxWidth(60).Sprintln(line)
		}
    title := pterm.LightRed(os.Args[1])
		pterm.DefaultBox.WithTitle(title).Println(wrappedText)

	}
}

func getCommonCommand(YamlInput inputmodels.YamlInput) (command.OSCommand, error) {
	cb := common.CommonCommandBuilder{}
	oscmd, err := cb.Build(YamlInput)
	return oscmd, err
}

func Dispatch(yamlInput inputmodels.YamlInput) error {
	config := os.Getenv("AOS")
	oscmd, err := getCommonCommand(yamlInput)
	if err != nil {
		return err
	}
	if config == "" {
		she := shell.ShellExecutor{}
		return she.Execute(oscmd)
	}

	for _, char := range config {
		switch char {
		case 'h':
			usage.PrintUsage()
		case 'H':
			printAOSDescription(yamlInput)
		case 'R':
			ce := console.ConsoleExecutor{}
			return ce.Execute(oscmd)
    case 'E':
			ee := explainer.ExplainerExecutor{YamlInput: yamlInput}
			return ee.Execute(oscmd)

		default:
		}

	}
	return nil
}
