package hw09structvalidator

import (
	"regexp"
	"strconv"
	"strings"
)

func validateString(value string, rules validationRules) error {
	for ruleName, rule := range rules {
		switch ruleName {
		case "in":
			valueSet := strings.Split(rule, ",")
			if !inValidation[string](value, valueSet) {
				return ErrNotAllowedValue
			}
		case "len":
			controlValue, err := strconv.Atoi(rule)
			if err != nil {
				return ErrIncorrectValidationRule
			}

			if !lenStringValidation(value, controlValue) {
				return ErrLenString
			}
		case "regexp":
			res, err := regExpStringValidation(value, rule)
			if err != nil {
				return ErrIncorrectValidationRule
			}

			if !res {
				return ErrNotMatchPattern
			}
		}
	}

	return nil
}

func lenStringValidation(value string, lValue int) bool {
	sl := []rune(value)

	return len(sl) == lValue
}

func regExpStringValidation(value, pattern string) (bool, error) {
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		return false, err
	}

	return matched, nil
}
