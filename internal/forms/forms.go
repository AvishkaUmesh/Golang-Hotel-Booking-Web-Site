package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form struct embeds url.Values and adds a errors field.
type Form struct {
	url.Values
	Errors errors
}

// NewForm initializes a custom Form struct.
func NewForm(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks that specific fields in a form submission are
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Required checks that specific fields in a form submission are
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks that a specific field in a form submission meets a minimum length requirement.
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
