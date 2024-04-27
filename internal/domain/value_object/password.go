package value_object

import (
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordMinLength = 8
	passwordMaxLength = 100
)

var (
	ErrInvalidPassword = errors.New("password must be longer than 8 characters, contain numbers and lowercase and uppercase letters")
)

type Password struct {
	Value string
}

func NewPassword(value string) (*Password, error) {
	password := Password{Value: value}
	err := password.validate()
	if err != nil {
		return nil, err
	}

	err = password.HashPassword(password.Value)
	if err != nil {
		return nil, err
	}

	return &password, nil
}

func (p *Password) validate() error {
	if len(p.Value) <= passwordMinLength || len(p.Value) > passwordMaxLength {
		return ErrInvalidPassword
	}
	hasLower := regexp.MustCompile(`[a-z]`).MatchString
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString

	if !hasLower(p.Value) || !hasUpper(p.Value) {
		return ErrInvalidPassword
	}
	return nil
}

func (p *Password) HashPassword(rawPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.Value = string(hash)
	return nil
}
