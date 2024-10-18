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

	for _, curRune := range runes {
		switch {
		case unicode.IsLetter(curRune):
			if escape {
				return "", ErrInvalidString
			}
			if prevRune == 0 {
				prevRune = curRune
			} else {
				result.WriteRune(prevRune)
				prevRune = curRune
			}
		case unicode.IsDigit(curRune):
			if prevRune == 0 && !escape {
				return "", ErrInvalidString
			}
			if escape {
				result.WriteRune(prevRune)
				prevRune = curRune
				escape = false
			} else {
				repCount, _ := strconv.Atoi(string(curRune))
				repeatedStr := strings.Repeat(string(prevRune), repCount)
				result.WriteString(repeatedStr)
				prevRune = 0
			}
		case curRune == '\\':
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
	}

	if prevRune != 0 {
		result.WriteRune(prevRune)
	}

	return result.String(), nil
}
