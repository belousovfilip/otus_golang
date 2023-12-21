package main

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
		Age    int      `validate:"min:18|max:50"`
		Email  string   `validate:"regexp:^\\w+@\\w+\\.\\w+$"`
		Role   UserRole `validate:"in:admin,stuff"`
		Phones []string `validate:"len:11"`
		meta   json.RawMessage
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
			in:          App{Version: "Equal"},
			expectedErr: ValidationErrors{},
		},
		{
			in: App{Version: "Few"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Version", Err: NewLenRule("Few", "5").Error()},
			},
		},
		{
			in: App{Version: "A lot of"},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Version", Err: NewLenRule("A lot of", "5").Error()},
			},
		},
		{
			in: User{
				ID:     strings.Repeat("5", 36),
				Name:   "Ivan",
				Age:    33,
				Email:  "test@test.com",
				Phones: strings.Split("1,2,3,4,5,6,7,8,9,10,11", ","),
				meta:   make(json.RawMessage, 0),
				Role:   UserRole("admin"),
			},
			expectedErr: ValidationErrors{},
		},
		{
			in: User{
				ID:     strings.Repeat("5", 36),
				Name:   "Ivan",
				Age:    33,
				Email:  "test@test.com",
				Phones: strings.Split("1,2,3,4,5,6,7,8,9,10,11", ","),
				meta:   make(json.RawMessage, 0),
				Role:   UserRole("_INVALID_"),
			},
			expectedErr: ValidationErrors{
				ValidationError{Field: "Role", Err: NewInRule("_INVALID_", []string{"admin", "stuff"}).Error()},
			},
		},
		{
			in: Response{
				Code: 200,
				Body: "<h1></h1>",
			},
			expectedErr: ValidationErrors{},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			tt := tt
			t.Parallel()
			err := Validate(tt.in)
			require.Equal(t, tt.expectedErr, err)
		})
	}
}
