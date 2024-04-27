package entity

import (
	"errors"
	valueobject "github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	"testing"
)

func TestCreateTransfer(t *testing.T) {
	type testcase struct {
		Value         float64
		Payer         *Account
		Payee         *Account
		ExpectedError error
	}

	t.Run("test should create new transfer without error", func(t *testing.T) {
		PayerUser, _ := createUser("André Luiz", "andre@gmail.com", "207.275.320-14")
		PayeeUser, _ := createUser("Marcos Silva", "marcos@gmail.com", "209.201.320-15")
		PayerAccount, _ := newAccount(*valueobject.NewID(), PayerUser, 100.0)
		PayeeAccount, _ := newAccount(*valueobject.NewID(), PayeeUser, 1000.0)

		testCases := []testcase{
			{Value: 100.0, Payer: PayerAccount, Payee: PayeeAccount, ExpectedError: nil},
		}
		for _, item := range testCases {
			_, err := NewTransfer(item.Value, item.Payer, item.Payee)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})

	t.Run("test should return invalid payer", func(t *testing.T) {
		PayerUser, _ := createUser("André Luiz", "andre@gmail.com", "207.275.320-14")
		userTypeId := valueobject.NewUserType(2)
		PayerUser.UserType = userTypeId
		PayeeUser, _ := createUser("Marcos Silva", "marcos@gmail.com", "209.201.320-15")
		PayerAccount, _ := newAccount(*valueobject.NewID(), PayerUser, 100.0)
		PayeeAccount, _ := newAccount(*valueobject.NewID(), PayeeUser, 1000.0)

		testCases := []testcase{
			{Value: 100.0, Payer: PayerAccount, Payee: PayeeAccount, ExpectedError: ErrInvalidPayer},
		}
		for _, item := range testCases {
			_, err := NewTransfer(item.Value, item.Payer, item.Payee)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})

	t.Run("test should return insufficient balance when try debit account", func(t *testing.T) {
		PayerUser, _ := createUser("André Luiz", "andre@gmail.com", "207.275.320-14")
		userTypeId := valueobject.NewUserType(1)
		PayerUser.UserType = userTypeId
		PayeeUser, _ := createUser("Marcos Silva", "marcos@gmail.com", "209.201.320-15")
		PayerAccount, _ := newAccount(*valueobject.NewID(), PayerUser, 10.0)
		PayeeAccount, _ := newAccount(*valueobject.NewID(), PayeeUser, 1000.0)

		testCases := []testcase{
			{Value: 100.0, Payer: PayerAccount, Payee: PayeeAccount, ExpectedError: ErrInsufficientBalance},
		}
		for _, item := range testCases {
			_, err := NewTransfer(item.Value, item.Payer, item.Payee)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})

	t.Run("test should return insufficient balance when try credit account", func(t *testing.T) {
		PayerUser, _ := createUser("André Luiz", "andre@gmail.com", "207.275.320-14")
		userTypeId := valueobject.NewUserType(1)
		PayerUser.UserType = userTypeId
		PayeeUser, _ := createUser("Marcos Silva", "marcos@gmail.com", "209.201.320-15")
		PayerAccount, _ := newAccount(*valueobject.NewID(), PayerUser, 10.0)
		PayeeAccount, _ := newAccount(*valueobject.NewID(), PayeeUser, 1000.0)

		testCases := []testcase{
			{Value: 0, Payer: PayerAccount, Payee: PayeeAccount, ExpectedError: ErrCreditValue},
		}
		for _, item := range testCases {
			_, err := NewTransfer(item.Value, item.Payer, item.Payee)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})
}

func createUser(userName, userEmail, userCpf string) (*User, error) {
	user, _ := CreateUserFactory(
		nil,
		"John Doe",
		"john.doe@example.com",
		"Password123A@s",
		"088.988.888-52",
		nil,
		1,
	)

	return user, nil
}
