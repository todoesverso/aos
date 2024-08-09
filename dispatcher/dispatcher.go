package dispatcher

import (
	"fmt"
	"os"
	"strings"

	"github.com/todoesverso/aos/command/builders/common"
	command "github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/executors/console"
	"github.com/todoesverso/aos/executors/explainer"
	"github.com/todoesverso/aos/executors/shell"
	inputmodels "github.com/todoesverso/aos/inputs/models"
	"github.com/todoesverso/aos/usage"

	"github.com/pterm/pterm"
)

type ValidEnvOptionsEnum byte

const (
	Charh ValidEnvOptionsEnum = 'h'
	CharH ValidEnvOptionsEnum = 'H'
	CharR ValidEnvOptionsEnum = 'R'
	CharE ValidEnvOptionsEnum = 'E'
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
		case rune(Charh):
			usage.PrintUsage()
		case rune(CharH):
			printAOSDescription(yamlInput)
		case rune(CharR):
			ce := console.ConsoleExecutor{}
			return ce.Execute(oscmd)
		case rune(CharE):
			ee := explainer.ExplainerExecutor{YamlInput: yamlInput}
			return ee.Execute(oscmd)
		default:
		}

	}
	return nil
}
