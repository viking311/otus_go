package hw09structvalidator

import (
	"bytes"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type validationRules map[string]string

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	if len(v) < 1 {
		return ""
	}
	buf := bytes.NewBufferString("")
	for i := 0; i < len(v); i++ {
		msg := fmt.Sprintf("Field: %s; Error: %s",
			v[i].Field, v[i].Err)
		buf.WriteString(msg)
		buf.WriteString("\n")
	}
	return buf.String()
}

var (
	UnSupportedTypeError    = errors.New("unsupported type")
	NotAllowedValueError    = errors.New("value is not in the allowed set")
	MinError                = errors.New("too small value")
	MaxError                = errors.New("too big value")
	LenStringError          = errors.New("incorrect string length")
	IncorrectValidationRule = errors.New("incorrect validation rule")
	NotMatchPatternError    = errors.New("value does not match pattern")
)

func Validate(v interface{}) error {
	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Struct {
		return UnSupportedTypeError
	}

	validationErrors := make(ValidationErrors, 0)

	for i := 0; i < value.NumField(); i++ {
		field := value.Type().Field(i)

		validateTag, ok := field.Tag.Lookup("validate")
		if !ok {
			continue
		}
		validationRules := parseTag(validateTag)

		switch field.Type.Kind() {
		case reflect.String:
			err := validateString(value.Field(i).String(), validationRules)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		case reflect.Int:
			err := validateInt(value.Field(i).Int(), validationRules)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		case reflect.Slice:
			err := sliceValidate(value.Field(i), validationRules)
			if err != nil {
				validationErrors = append(validationErrors, ValidationError{
					Field: field.Name,
					Err:   err,
				})
			}
		default:
			continue
		}
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func parseTag(tagValue string) validationRules {
	rules := make(validationRules)

	ruleStrings := strings.Split(tagValue, "|")
	for _, r := range ruleStrings {
		rSplit := strings.Split(r, ":")
		if len(rSplit) != 2 {
			continue
		}
		rules[rSplit[0]] = rSplit[1]
	}

	return rules
}
