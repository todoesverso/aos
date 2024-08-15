package argpars

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/todoesverso/aos/inputs/models"
)

func TestParseArgument_RawArgument(t *testing.T) {
	argument := models.Argument{
		Raw: "ffmpeg -i input.mp4 -codec copy output.mp4",
	}

	expected := []string{"ffmpeg", "-i", "input.mp4", "-codec", "copy", "output.mp4"}
	result := ParseArgument(argument)

	assert.Equal(t, expected, result, "Expected parsed arguments to match the split raw argument")
}

func TestParseArgument_OptionAndValue(t *testing.T) {
	argument := models.Argument{
		Option: "-i",
		Value:  "input.mp4",
	}

	expected := []string{"-i", "input.mp4"}
	result := ParseArgument(argument)

	assert.Equal(t, expected, result, "Expected parsed arguments to include both option and value")
}

func TestParseArgument_EmptyFields(t *testing.T) {
	argument := models.Argument{}

	expected := []string{}
	result := ParseArgument(argument)

	assert.Equal(t, expected, result, "Expected parsed arguments to be empty for empty fields")
}

func TestParseArgument_OnlyOption(t *testing.T) {
	argument := models.Argument{
		Option: "-v",
	}

	expected := []string{"-v"}
	result := ParseArgument(argument)

	assert.Equal(t, expected, result, "Expected parsed arguments to include only the option")
}

func TestParseArgument_OnlyValue(t *testing.T) {
	argument := models.Argument{
		Value: "output.mp4",
	}

	expected := []string{"output.mp4"}
	result := ParseArgument(argument)

	assert.Equal(t, expected, result, "Expected parsed arguments to include only the value")
}
