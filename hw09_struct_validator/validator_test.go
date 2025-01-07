package hw09structvalidator

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

type UserRole string

// Test the function on different structures and other types.
type (
	User struct {
		ID     string `json:"id" validate:"len:36"`
		Name   string
		Age    int             `validate:"min:18|max:50"`
		Email  string          `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole        `validate:"in:admin,stuff"`
		Phones []string        `validate:"len:11"`
		meta   json.RawMessage //nolint:unused
	}

	App struct {
		Version string `validate:"len:5"`
	}

	Token struct {
		Header    []byte
		Payload   []byte
		Signature []byte
	}

	Response struct {
		Code int    `validate:"in:200,404,500"`
		Body string `json:"omitempty"`
	}
)

func TestValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in:          UserRole("admin"),
			expectedErr: ErrUnsupportedType,
		},
		{
			in: Response{
				Code: 200,
				Body: "body",
			},
			expectedErr: nil,
		},
		{
			in: Response{
				Code: 201,
				Body: "body",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Code",
					Err:   ErrNotAllowedValue,
				},
			},
		},
		{
			in: App{
				Version: "12345",
			},
			expectedErr: nil,
		},
		{
			in: App{
				Version: "1234",
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "Version",
					Err:   ErrLenString,
				},
			},
		},
		{
			in: Token{
				Header:    []byte("hello"),
				Payload:   []byte("world"),
				Signature: []byte("!!!!"),
			},
			expectedErr: nil,
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
			_ = tt
		})
	}
}

func TestUserValidate(t *testing.T) {
	tests := []struct {
		in          interface{}
		expectedErr error
	}{
		{
			in: User{
				ID:    strings.Repeat("*", 36),
				Age:   19,
				Email: "120@ya.ru",
				Role:  "admin",
				Phones: []string{
					"+7812345678",
				},
			},
			expectedErr: nil,
		},
		{
			in: User{
				ID:    strings.Repeat("*", 3),
				Age:   15,
				Email: "120@ya",
				Role:  "user",
				Phones: []string{
					"+7812345678",
					"+7812345",
				},
			},
			expectedErr: ValidationErrors{
				ValidationError{
					Field: "ID",
					Err:   ErrLenString,
				},
				ValidationError{
					Field: "Age",
					Err:   ErrMin,
				},
				ValidationError{
					Field: "Email",
					Err:   ErrNotMatchPattern,
				},
				ValidationError{
					Field: "Role",
					Err:   ErrNotAllowedValue,
				},
				ValidationError{
					Field: "Phones",
					Err:   ErrLenString,
				},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
			_ = tt
		})
	}
}
