package entity

import (
	"errors"
)

const (
	invalidPayer = 2
)

type Transfer struct {
	id    *int64
	value float64
	payer *Account
	payee *Account
}

var (
	ErrInvalidPayer = errors.New("shopkeeper can't make a transfer")
)

func NewTransfer(id *int64, value float64, payer, payee *Account) (*Transfer, error) {
	transfer := &Transfer{
		id:    id,
		value: value,
		payer: payer,
		payee: payee,
	}

	err := transfer.validate()
	if err != nil {
		return nil, err
	}
	return transfer, nil
}

func (t *Transfer) validate() error {
	if t.isInvalidPayer() {
		return ErrInvalidPayer
	}
	return t.makeTransfer()
}

func (t *Transfer) isInvalidPayer() bool {
	return t.payer.User().UserType.Value == invalidPayer
}

func (t *Transfer) makeTransfer() error {
	err := t.payer.DebitAccount(t.value)
	if err != nil {
		return err
	}
	err = t.payee.CreditAccount(t.value)
	if err != nil {
		return err
	}
	return nil
}

func (t *Transfer) Value() float64 {
	return t.value
}

func (t *Transfer) Payer() *int64 {
	return t.payer.ID()
}

func (t *Transfer) Payee() *int64 {
	return t.payee.ID()
}

func (t *Transfer) PayeeBalance() float64 {
	return t.payee.Balance()
}
