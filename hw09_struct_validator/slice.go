package hw09structvalidator

import "reflect"

func sliceValidate(value reflect.Value, rules validationRules) error {
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		switch elem.Kind() {
		case reflect.String:
			err := validateString(elem.String(), rules)
			if err != nil {
				return err
			}
		case reflect.Int:
			err := validateInt(elem.Int(), rules)
			if err != nil {
				return err
			}
		default:
		}
	}

	return nil
}
