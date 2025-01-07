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
				return NotAllowedValueError
			}
		case "len":
			controlValue, err := strconv.Atoi(rule)
			if err != nil {
				return IncorrectValidationRule
			}

			if !lenStringValidation(value, controlValue) {
				return LenStringError
			}
		case "regexp":
			res, err := regExpStringValidation(value, rule)
			if err != nil {
				return IncorrectValidationRule
			}

			if !res {
				return NotMatchPatternError
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
		return false, nil
	}

	return matched, nil
}
