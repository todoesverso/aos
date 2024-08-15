package main

import (
	"fmt"
	"os"

	"github.com/todoesverso/aos/command/builders/common"
	"github.com/todoesverso/aos/dispatcher"
	"github.com/todoesverso/aos/executors/usage"
	"github.com/todoesverso/aos/inputs/file"
)

func run(inputFile string) error {
	if len(os.Args) < common.DEFAULT_ARGS_NUMBER {
		usage.PrintUsage()
		return fmt.Errorf("insufficient arguments")
	}

	yamlData, err := file.ParseInputFile(inputFile)
	if err != nil {
		return fmt.Errorf("error parsing input file: %v", err)
	}

	err = dispatcher.Dispatch(yamlData)
	if err != nil {
		return fmt.Errorf("error dispatching: %v", err)
	}
	return nil
}

func main() {
	if len(os.Args) < common.DEFAULT_ARGS_NUMBER {
		usage.PrintUsage()
		os.Exit(1)
	}

	inputFile := os.Args[1]
	err := run(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
