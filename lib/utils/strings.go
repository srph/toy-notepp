package utils

import (
	"strings"
)

// Remove duplicated strings from the array
func StringUnique(strings []string) []string {
	var validated []string

	for _, str := range strings {
		if !ArrayContains(validated, str) {
			validated = append(validated, str)
		}
	}

	return validated
}

// Trim the the specified delimiter on
// the start and end of each string.
func TrimStrings(strs []string, delimiter string) []string {
	var trimmed []string

	for _, str := range strs {
		trimmed = append(trimmed, strings.Trim(str, delimiter))
	}

	return trimmed
}