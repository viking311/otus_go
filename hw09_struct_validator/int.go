package hw09structvalidator

import (
	"strconv"
	"strings"
)

func validateInt(value int64, rules validationRules) error {
	for ruleName, rule := range rules {
		switch ruleName {
		case "in":
			rawValueSet := strings.Split(rule, ",")
			valueSet := make([]int64, 0)
			for _, rawItem := range rawValueSet {
				castValue, err := strconv.Atoi(rawItem)
				if err != nil {
					return ErrIncorrectValidationRule
				}

				valueSet = append(valueSet, int64(castValue))
			}

			if !inValidation[int64](value, valueSet) {
				return ErrNotAllowedValue
			}
		case "min":
			controlValue, err := strconv.Atoi(rule)
			if err != nil {
				return ErrIncorrectValidationRule
			}

			if !intMinValidation(value, int64(controlValue)) {
				return ErrMin
			}
		case "max":
			controlValue, err := strconv.Atoi(rule)
			if err != nil {
				return ErrIncorrectValidationRule
			}

			if !intMaxValidation(value, int64(controlValue)) {
				return ErrMax
			}
		}
	}
	return nil
}

func intMinValidation(value, limit int64) bool {
	return value >= limit
}

func intMaxValidation(value, limit int64) bool {
	return value <= limit
}
