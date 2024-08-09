package file

import (
	"fmt"
	"os"

	"github.com/todoesverso/aos/inputs/models"
	"gopkg.in/yaml.v3"
)

func ParseInputFile(inputFile string) (models.YamlInput, error) {
	yamlData := &models.YamlInput{}

	yamlFile, err := os.ReadFile(inputFile)
	if err != nil {
		return *yamlData, fmt.Errorf("failed to read file %s: %w", inputFile, err)
	}
	err = yaml.Unmarshal(yamlFile, yamlData)
	if err != nil {
		return *yamlData, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return *yamlData, nil
}
