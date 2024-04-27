package entity

import (
	"errors"

	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

const (
	minBalance = 0
)

var (
	ErrInsufficientBalance              = errors.New("insufficient balance")
	ErrCreditValue                      = errors.New("creditValue can't be zero")
	ErrCreateAccountWithNegativeBalance = errors.New("it is not possible to create an account with a negative balance")
)

type Account struct {
	id      value_object.ID
	user    *User
	balance float64
}

func newAccount(id value_object.ID, user *User, balance float64) (*Account, error) {
	account := &Account{
		id:      id,
		user:    user,
		balance: balance,
	}

	err := account.validate()
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (acc *Account) validate() error {
	if acc.balance < minBalance {
		return ErrCreateAccountWithNegativeBalance
	}
	return nil
}

func (acc *Account) CreditAccount(value float64) error {
	if value == 0 {
		return ErrCreditValue
	}
	acc.balance += value
	return nil
}

func (acc *Account) DebitAccount(value float64) error {
	if acc.isInsufficientBalance() {
		return ErrInsufficientBalance
	}
	if acc.balance-value < minBalance {
		return ErrInsufficientBalance
	}
	acc.balance -= value
	return nil
}

func (acc *Account) isInsufficientBalance() bool {
	return acc.balance <= minBalance
}

func (acc *Account) ID() value_object.ID {
	return acc.id
}

func (acc *Account) User() *User {
	return acc.user
}

func (acc *Account) Balance() float64 {
	return acc.balance
}
