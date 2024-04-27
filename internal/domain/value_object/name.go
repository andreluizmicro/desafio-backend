package value_object

import (
	"errors"
	"strings"
)

const (
	nameMinLength = 3
	nameMaxLength = 200
)

var (
	ErrInvalidName = errors.New("the name must be between 3 and 200 characters")
)

type Name struct {
	Value string
}

func NewName(value string) (*Name, error) {
	name := Name{Value: strings.ToUpper(value)}
	err := name.validate()
	if err != nil {
		return nil, err
	}
	return &name, nil
}

func (n *Name) validate() error {
	if len(n.Value) <= nameMinLength || len(n.Value) > nameMaxLength {
		return ErrInvalidName
	}
	return nil
}
