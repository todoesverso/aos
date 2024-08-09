package usage

import (
	"fmt"
	"os"
)

const ENV_VAR = "AOS"

func PrintUsage() {
  exec_name := os.Args[0]
	fmt.Printf("AliasOnSteroids\n\n")
	fmt.Printf("Usage:\n")
	fmt.Printf("\t%s <alias.yaml> [positional arguments of the alias]\n", exec_name)

	fmt.Printf("Options:\n")
	fmt.Printf("\tIn order to keep CLI arguments as straightforward as possible,\n")
	fmt.Printf("\toptions are passed thru the %s environment variables.\n\n", ENV_VAR)
	fmt.Printf("\t%s=h %s <alias.yaml>\t# Prints this usage\n", ENV_VAR, exec_name)
	fmt.Printf("\t%s=H %s <alias.yaml>\t# Prints a helper short description of the alias\n", ENV_VAR, exec_name)
	fmt.Printf("\t%s=E %s <alias.yaml>\t# Prints a helper long description of the alias\n", ENV_VAR, exec_name)
	fmt.Printf("\t%s=R %s <alias.yaml>\t# Renders the command and prints it to stdout\n", ENV_VAR, exec_name)
}
