package explainer

import (
	"strings"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/todoesverso/aos/command/models"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

type ExplainerExecutor struct {
	YamlInput inputmodels.YamlInput
}

func wrapTextWithTree(text string, maxWidth, level int) string {
	if len(text) < maxWidth {
		return text
	}
	wrappedText := pterm.DefaultParagraph.WithMaxWidth(maxWidth).Sprintln(text)
	lines := strings.Split(wrappedText, "\n")
	prefix := pterm.ThemeDefault.TreeStyle.Sprint(strings.Repeat("│   ", level))

	// Apply prefix to the wrapped lines, keeping the first line unchanged
	for i := range lines {
		if i > 0 {
			lines[i] = prefix + lines[i]
		}
	}

	// Join the lines back together
	return strings.Join(lines, "\n")
}

func (y ExplainerExecutor) Execute(cmd models.OSCommand) error {
	lightGreenBoldStyle := pterm.NewStyle(pterm.FgLightGreen, pterm.BgDefault, pterm.Bold)
	leveledList := pterm.LeveledList{}
	for _, c := range y.YamlInput.Arguments {
		var text string
		if c.Positional != nil {
			text = lightGreenBoldStyle.Sprint(c.Positional.Name) + ": "
			text += wrapTextWithTree(c.Positional.Description, 60, 1)
		} else {
			text = lightGreenBoldStyle.Sprint(c.Option )+ ": "
			text += wrapTextWithTree(c.Description, 60, 1)
		}

		leveledList = append(leveledList, pterm.LeveledListItem{Level: 0, Text: text})
	}

	lightRedBoldStyle := pterm.NewStyle(pterm.FgLightRed, pterm.BgDefault, pterm.Bold)
	text := lightRedBoldStyle.Sprint(y.YamlInput.Command.Exec) + ": "
	text += wrapTextWithTree(y.YamlInput.Command.Description, 60, 0)
	root := putils.TreeFromLeveledList(leveledList)
	root.Text = text

	pterm.DefaultTree.WithRoot(root).Render()
	return nil
}
