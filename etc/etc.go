package etc

import (
	"strings"
)

// converts text to lowercase substrings
func StringSlice(m string) []string {

	m = strings.ToLower(m)
	return strings.Split(m, " ")
}

// centers text
// im doing this the shitty not expandable way because ive been defeated
func ToCenter(s string) string {
	switch len(s) {
	case 1:
		return " " + s + " "
	case 2:
		return " " + s
	default:
		return s
	}
}

// adds 'i' amount of 'x' to string 's'
func stringLoop(s string, x string, i int) string {

	for len(s) < i {
		s += x
	}
	return s
}

// centers text, properly, but for some reason throws a hissy fit if i use
// spaces
// TODO: Fix this
func ToCenterAlt(s string, i int) string {
	if i > len(s) {
		o := i - len(s)
		ns := stringLoop("", " ", o) + s

		return ns
	}
	return s
}
