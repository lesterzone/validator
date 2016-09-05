package validator

import "regexp"

// Valid struct based
type Validator struct {
	Valid  bool
	Errors map[string]string
}

// Validator instance
func I() Validator {
	return Validator{Errors: make(map[string]string), Valid: true}
}

// Present validate presence of a field
func (v *Validator) Present(field string, value string) bool {
	if value == "" {
		v.Valid = false
		v.Errors[field] = "Required"
	}

	return v.Valid
}

// EmailFormat validate presence of a field
// http://stackoverflow.com/questions/46155/validate-email-address-in-javascript
func (v *Validator) EmailFormat(field string, value string) bool {
	re := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

	if !re.MatchString(value) {
		v.Valid = false
		v.Errors[field] = "Invalid email format"
	}

	return v.Valid
}

// MinLenght validate min length against provided value
func (v *Validator) MinLenght(field string, value string, min int) bool {

	if len(value) < min {
		v.Valid = false
		v.Errors[field] = "min length required is " + string(min)
	}

	return v.Valid
}

// IsValid to check if we have any errors arrociated to v
func (v *Validator) IsValid() bool {
	return !v.Valid
}
