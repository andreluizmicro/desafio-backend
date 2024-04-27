package value_object

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email struct {
	Value string
}

func NewEmail(value string) (*Email, error) {
	email := Email{Value: value}
	err := email.validate()
	if err != nil {
		return nil, err
	}
	return &email, nil
}

func (e *Email) validate() error {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(e.Value) {
		return ErrInvalidEmail
	}
	return nil
}
