package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var result strings.Builder
	var prevRune rune
	var escape bool
	result.Reset()
	for _, curRune := range input {
		switch {
		case unicode.IsLetter(curRune):
			if escape {
				return "", ErrInvalidString
			}
			if prevRune == 0 {
				prevRune = curRune
			} else {
				writeRune(prevRune, 1, &result)
				prevRune = curRune
			}
		case unicode.IsDigit(curRune):
			if prevRune == 0 && !escape {
				return "", ErrInvalidString
			}
			if escape {
				if prevRune != 0 {
					writeRune(prevRune, 1, &result)
				}
				prevRune = curRune
				escape = false
			} else {
				repCount, _ := strconv.Atoi(string(curRune))
				writeRune(prevRune, repCount, &result)
				prevRune = 0
			}
		case curRune == '\\':
			if escape {
				escape = false
				writeRune(prevRune, 1, &result)
				prevRune = '\\'
			} else {
				escape = true
			}
		default:
			return "", ErrInvalidString
		}
	}

	writeRune(prevRune, 1, &result)

	return result.String(), nil
}

func writeRune(r rune, count int, result *strings.Builder) {
	if r == 0 || count == 0 {
		return
	}

	str := string(r)
	if count > 1 {
		str = strings.Repeat(str, count)
	}

	result.WriteString(str)
}
