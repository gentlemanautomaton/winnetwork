package netresource

import (
	"fmt"
	"strings"
)

// MakeDriveLetter sanitizes and validates the given value as a drive letter
// between A and Z. If a trailing ":" character is provided in the input, it
// is removed from the result.
func MakeDriveLetter(value string) (letter string, err error) {
	letter = strings.ToUpper(strings.TrimSuffix(value, ":"))
	if len(letter) != 1 || letter[0] < 'A' || letter[0] > 'Z' {
		return "", fmt.Errorf("invalid drive letter: %s", value)
	}
	return letter, nil
}
