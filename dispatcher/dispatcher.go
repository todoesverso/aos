package dispatcher

import (
	"errors"
	"os"

	"github.com/todoesverso/aos/command/builders/common"
	command "github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/executors"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

func getCommonCommand(yamlInput inputmodels.YamlInput) (command.OSCommand, error) {
	return common.CommonCommandBuilder{}.Build(yamlInput)
}

func Dispatch(yamlInput inputmodels.YamlInput) error {
	config := os.Getenv("AOS")
	oscmd, err := getCommonCommand(yamlInput)

	if err != nil {
		return err
	}

	executorRegistry := executors.InitExecutors(yamlInput)
	if config == "" {
		if defaultExecutor, ok := executorRegistry.GetExecutor(executors.Chard); ok {
			return defaultExecutor.Execute(oscmd)
		}
		return errors.New("Executor not found")
	}

	for _, char := range config {
		if exec, exists := executorRegistry.GetExecutor(executors.EnvOption(char)); exists {
			return exec.Execute(oscmd)
		}
	}

	return nil
}
