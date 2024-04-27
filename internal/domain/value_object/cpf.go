package value_object

import (
	"errors"
	"regexp"
)

const cpfLength = 11

var (
	ErrInvalidCPF = errors.New("invalid cpf")
)

type CPF struct {
	Value string
}

func NewCPF(value string) (*CPF, error) {
	cpf := CPF{Value: value}
	err := cpf.validate()
	if err != nil {
		return nil, err
	}
	return &cpf, nil
}

func (c *CPF) validate() error {
	cpfRegex := regexp.MustCompile(`[^0-9]`)
	cpf := cpfRegex.ReplaceAllString(c.Value, "")

	if len(cpf) != cpfLength {
		return ErrInvalidCPF
	}
	return nil
}
