package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	runes := []rune(input)

	var result strings.Builder
	var prevRune rune
	var escape bool
	var i int

	for i < len(runes) {
		switch {
		case unicode.IsLetter(runes[i]):
			if escape {
				return "", ErrInvalidString
			}
			if prevRune == 0 {
				prevRune = runes[i]
			} else {
				result.WriteRune(prevRune)
				prevRune = runes[i]
			}
		case unicode.IsDigit(runes[i]):
			if prevRune == 0 && !escape {
				return "", ErrInvalidString
			}
			if escape {
				result.WriteRune(prevRune)
				prevRune = runes[i]
				escape = false
			} else {
				repCount, _ := strconv.Atoi(string(runes[i]))
				repeatedStr := strings.Repeat(string(prevRune), repCount)
				result.WriteString(repeatedStr)
				prevRune = 0
			}
		case runes[i] == '\\':
			if escape {
				escape = false
				result.WriteRune(prevRune)
				prevRune = '\\'
			} else {
				escape = true
			}
		default:
			return "", ErrInvalidString
		}

		i++
	}

	if prevRune != 0 {
		result.WriteRune(prevRune)
	}

	return result.String(), nil
}
