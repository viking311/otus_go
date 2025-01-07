package hw09structvalidator

type integer interface {
	int | int8 | int16 | int32 | int64 | byte
}

func inValidation[T integer | string](value T, controlSet []T) bool {
	for _, item := range controlSet {
		if value == item {
			return true
		}
	}

	return false
}
