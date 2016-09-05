package validator_test

import (
	"testing"

	"github.com/lesterzone/validator"
)

func TestValidator(t *testing.T) {
	// arrange & act
	v := validator.I()

	if !v.Valid {
		t.Errorf("expected instance to init with true, got: %v", v.Valid)
	}

	if v.Errors == nil {
		t.Errorf("expected instance to init with empty errors map, got: %v", v.Errors)
	}
}

type E struct {
	Email string
}

func (e *E) IsValid() (bool, map[string]string) {
	v := validator.I()
	if e.Email != "" {
		v.Errors["email"] = "email required"
		v.Valid = false
	}
	return v.Valid, v.Errors
}

func TestPresent(t *testing.T) {
	v := validator.I()

	b := v.Present("foo", "")

	if b {
		t.Errorf("expect Present to validate precence of field")
	}
}

func TestEmailFormat(t *testing.T) {
	// arrange & act
	var cases = []struct {
		in  string
		out bool
	}{
		{"foo", false},
		{"f@f.com", true},
		{"f@", false},
		{"@f.com", false},
	}

	for _, option := range cases {
		v := validator.I()
		valid := v.EmailFormat("email", option.in)

		if valid != option.out {
			t.Errorf("expect %v to be invalid email, %v, %v", option.in, valid, option.out)
		}
	}
}

func TestMinLength(t *testing.T) {
	// arrange & act
	var cases = []struct {
		in  string
		min int
		out bool
	}{
		{"foo", 4, false},
		{"foobar", 5, true},
		{"f@", 2, true},
		{"f", 1, true},
	}

	for _, option := range cases {
		v := validator.I()
		valid := v.MinLenght("property", option.in, option.min)

		if valid != option.out {
			t.Errorf("in %v expected: %v actual: %v", option.in, option.out, valid)
		}
	}
}

func TestIsValid(t *testing.T) {
	// arrage & act
	v := validator.I()
	b := v.IsValid()

	// assert
	if b {
		t.Errorf("expect Present to validate precence of field")
	}
}
