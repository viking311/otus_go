package hw09structvalidator

import "reflect"

func sliceValidate(value reflect.Value, rules validationRules) error {
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		switch {
		case elem.Kind() == reflect.String:
			err := validateString(elem.String(), rules)
			if err != nil {
				return err
			}
		case elem.Kind() == reflect.Int:
			err := validateInt(elem.Int(), rules)
			if err != nil {
				return err
			}
		default:
			return ErrUnsupportedType
		}
	}

	return nil
}
