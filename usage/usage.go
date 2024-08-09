package usage

import (
  "os"
  "fmt"
)

func PrintUsage() {
  fmt.Printf("AliasOnSteroids\n\n")
  fmt.Printf("Usage:\n")
  fmt.Printf("\t%s <alias.yaml> [positional arguments of the alias]\n", os.Args[0])

  fmt.Printf("Options:\n")
  fmt.Printf("\tIn order to keep CLI arguments as straightforward as possible,\n")
  fmt.Printf("\toptions are passed thru environment variables.\n\n")
  fmt.Printf("\tAOS=h %s <alias.yaml>\t# Prints this usage\n", os.Args[0])
  fmt.Printf("\tAOS=H %s <alias.yaml>\t# Prints a helper description of the alias\n", os.Args[0])
  fmt.Printf("\tAOS=r %s <alias.yaml>\t# Renders the command and prints it to stdout\n", os.Args[0])
}
