package executors

import (
	"testing"

	"github.com/stretchr/testify/assert"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

func TestNewExecutorRegistry(t *testing.T) {
	registry := NewExecutorRegistry()
	assert.NotNil(t, registry, "ExecutorRegistry should be initialized")
	assert.NotNil(t, registry.executors, "Executor map should be initialized")
}

func TestRegisterAndGetExecutor(t *testing.T) {
	registry := NewExecutorRegistry()
	yamlInput := inputmodels.YamlInput{}
	explainerExecutor := NewExplainerExecutor(yamlInput)

	registry.Register(CharE, explainerExecutor)

	retrievedExecutor, found := registry.GetExecutor(CharE)
	assert.True(t, found, "Executor should be found in the registry")
	assert.Equal(t, explainerExecutor, retrievedExecutor, "Retrieved executor should match the registered executor")
}

func TestGetExecutor_NotFound(t *testing.T) {
	registry := NewExecutorRegistry()

	_, found := registry.GetExecutor(CharH)
	assert.False(t, found, "Executor should not be found in the registry if it was not registered")
}

func TestInitExecutors(t *testing.T) {
	yamlInput := inputmodels.YamlInput{}
	registry := InitExecutors(yamlInput)

	// Test for all expected executors
	executorsToTest := []struct {
		envOption EnvOption
		expected  bool
	}{
		{CharR, true},
		{CharE, true},
		{Charh, true},
		{CharH, true},
		{CharX, true},
		{Chard, true},
	}

	for _, test := range executorsToTest {
		_, found := registry.GetExecutor(test.envOption)
		assert.Equal(t, test.expected, found, "Executor should be registered")
	}
}
