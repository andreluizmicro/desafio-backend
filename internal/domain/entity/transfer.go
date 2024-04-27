package entity

import "errors"

const (
	invalidPayer = 2
)

type Transfer struct {
	value float64
	payer *Account
	payee *Account
}

var (
	ErrInvalidPayer = errors.New("shopkeeper can't make a transfer")
)

func NewTransfer(value float64, payer, payee *Account) (*Transfer, error) {
	transfer := &Transfer{
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
