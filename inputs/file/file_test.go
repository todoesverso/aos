package file

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseInputFile_Success(t *testing.T) {
	// Mock input YAML content
	yamlContent := `
command:
  exec: "ffmpeg"
  description: "A command to run ffmpeg"
arguments:
  - option: "-i"
    value: "input.mp4"
`

	// Create a temporary file with the mock YAML content
	tmpFile, err := os.CreateTemp("", "test-input-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(yamlContent))
	assert.NoError(t, err)
	tmpFile.Close()

	// Call the function to parse the file
	parsedYaml, err := ParseInputFile(tmpFile.Name())

	// Validate the parsed data
	assert.NoError(t, err)
	assert.Equal(t, "ffmpeg", parsedYaml.Command.Exec)
	assert.Equal(t, "A command to run ffmpeg", parsedYaml.Command.Description)
	assert.Len(t, parsedYaml.Arguments, 1)
	assert.Equal(t, "-i", parsedYaml.Arguments[0].Option)
	assert.Equal(t, "input.mp4", parsedYaml.Arguments[0].Value)
}

func TestParseInputFile_FileNotFound(t *testing.T) {
	// Call the function with a non-existent file
	_, err := ParseInputFile("nonexistent.yaml")

	// Validate that an error is returned
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to read file")
}

func TestParseInputFile_InvalidYAML(t *testing.T) {
	// Mock invalid YAML content
	invalidYamlContent := `
command
  exec: "ffmpeg"
  description: "A command to run ffmpeg"
`

	// Create a temporary file with the invalid YAML content
	tmpFile, err := os.CreateTemp("", "test-invalid-*.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(invalidYamlContent))
	assert.NoError(t, err)
	tmpFile.Close()

	// Call the function to parse the file
	_, err = ParseInputFile(tmpFile.Name())

	// Validate that an error is returned
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to unmarshal YAML")
}
