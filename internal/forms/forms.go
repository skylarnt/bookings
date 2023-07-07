package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type FormInfo struct {
	url.Values
	Errors errors
}

func New(data url.Values) *FormInfo {
	return &FormInfo{
		data,
		errors(map[string][]string{}),
	}
}

func (f *FormInfo) Valid() bool {
	return len(f.Errors) == 0
}

func (f *FormInfo) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This is a required field")
		}
	}

}

func (f *FormInfo) MinLength(field string, length int, r *http.Request) bool {
	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true

}

func (f *FormInfo) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x != "" {
		return true
	}
	f.Errors.Add(field, "This is a required field")
	return true
}

func (f *FormInfo) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
