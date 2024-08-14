package main

import (
	"fmt"
	"os"

	"github.com/todoesverso/aos/command/builders/common"
	"github.com/todoesverso/aos/dispatcher"
	"github.com/todoesverso/aos/executors/usage"
	"github.com/todoesverso/aos/inputs/file"
)

func main() {
	if len(os.Args) < common.DEFAULT_ARGS_NUMBER {
		usage.PrintUsage()
		os.Exit(1)
	}
	inputFile := os.Args[1]

	yamlData, err := file.ParseInputFile(inputFile)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	err = dispatcher.Dispatch(yamlData)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	os.Exit(0)
}
