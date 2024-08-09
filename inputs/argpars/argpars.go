package argpars

import (
	"strings"

	"github.com/todoesverso/aos/inputs/models"
)

func appendIfNotEmpty(slice []string, value string) []string {
	if value != "" {
		slice = append(slice, value)
	}
	return slice
}

func ParseArgument(argument models.Argument) []string {
	var ret []string

	if argument.Raw != "" {
		arguments := strings.Split(argument.Raw, " ")
		ret = append(ret, arguments...)
		return ret
	}

	ret = appendIfNotEmpty(ret, argument.Option)
	ret = appendIfNotEmpty(ret, argument.Value)
	return ret
}
