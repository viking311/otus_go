package app

type FieldValidationError struct {
	field string
	msg   string
}

func (fle *FieldValidationError) Error() string {
	return "validation error: field " + fle.field + ": " + fle.msg
}
