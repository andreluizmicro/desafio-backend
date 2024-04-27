package value_object

import (
	"errors"
	"regexp"
)

const cnpjLength = 14

var (
	ErrInvalidCNPJ = errors.New("invalid cnpj")
)

type CNPJ struct {
	Value string
}

func NewCNPJ(value string) (*CNPJ, error) {
	cnpj := CNPJ{Value: value}
	err := cnpj.validate()
	if err != nil {
		return nil, err
	}
	return &cnpj, nil
}

func (c *CNPJ) validate() error {
	cpfRegex := regexp.MustCompile(`[^0-9]`)
	cpf := cpfRegex.ReplaceAllString(c.Value, "")

	if len(cpf) != cnpjLength {
		return ErrInvalidCNPJ
	}
	return nil
}
