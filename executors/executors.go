package executors

import (
	command "github.com/todoesverso/aos/command/models"
	"github.com/todoesverso/aos/executors/console"
	"github.com/todoesverso/aos/executors/explainer"
	"github.com/todoesverso/aos/executors/helper"
	"github.com/todoesverso/aos/executors/shell"
	"github.com/todoesverso/aos/executors/usage"
	inputmodels "github.com/todoesverso/aos/inputs/models"
)

type EnvOption byte

const (
	Charh EnvOption = 'h'
	CharH EnvOption = 'H'
	CharR EnvOption = 'R'
	CharX EnvOption = 'X'
	CharE EnvOption = 'E'
	Chard EnvOption = 'd' // DEFAULT EXECUTOR
)

// This function will init the ExecutorsRegistry and register all executors
func InitExecutors(YamlInput inputmodels.YamlInput) *ExecutorRegistry {
	er := NewExecutorRegistry()
	er.Register(CharR, console.ConsoleExecutor{})
	er.Register(CharE, explainer.ExplainerExecutor{YamlInput: YamlInput})
	er.Register(Charh, usage.UsageExecutor{})
	er.Register(CharH, helper.HelperExecutor{YamlInput: YamlInput})
	er.Register(CharX, shell.ShellExecutor{})
	er.Register(Chard, shell.ShellExecutor{})
	return er
}

type Executor interface {
	Execute(cmd command.OSCommand) error
}

type OutputProcessor interface {
	Executor
}

type ExecutorRegistry struct {
	executors map[EnvOption]Executor
}

func NewExecutorRegistry() *ExecutorRegistry {
	return &ExecutorRegistry{executors: make(map[EnvOption]Executor)}
}

func (r *ExecutorRegistry) Register(char EnvOption, executor Executor) {
	r.executors[char] = executor
}

func (r *ExecutorRegistry) GetExecutor(char EnvOption) (Executor, bool) {
	executor, found := r.executors[char]
	return executor, found
}

func NewExplainerExecutor(yamlInput inputmodels.YamlInput) Executor {
	return explainer.ExplainerExecutor{YamlInput: yamlInput}
}
