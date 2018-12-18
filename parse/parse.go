package parse

import (
	"regexp"

	"github.com/oct2pus/botutil/logging"
)

func CheckFormatted(input string, rgxp string) bool {
	// todo: fix +- bullshit with regexp
	compare, err := regexp.MatchString(rgxp, input)
	logging.CheckError(err)

	if compare {
		return true
	}
	return false
}
