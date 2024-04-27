package entity

import (
	"errors"
	valueobject "github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	"testing"
)

func TestCreateUser(t *testing.T) {
	type testcase struct {
		ID            *valueobject.ID
		Name          string
		Email         string
		Password      string
		CPF           string
		CNPJ          *string
		UserType      int
		ExpectedError error
	}

	t.Run("test success create user", func(t *testing.T) {
		validCnpj := "43.492.164/0001-96"
		invalidCnpj := "000000"
		testCases := []testcase{
			{
				ID:            nil,
				Name:          "André Silva",
				Email:         "andre@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      1,
				ExpectedError: nil,
			},
			{
				Name:          "aaa",
				Email:         "andre@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      1,
				ExpectedError: valueobject.ErrInvalidName,
			},
			{
				Name:          "Paulo Silva",
				Email:         "aaa",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      1,
				ExpectedError: valueobject.ErrInvalidEmail,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "a111aaa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      1,
				ExpectedError: valueobject.ErrInvalidPassword,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "a111aaaa@dsa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      1,
				ExpectedError: valueobject.ErrInvalidPassword,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458",
				CNPJ:          nil,
				UserType:      1,
				ExpectedError: valueobject.ErrInvalidCPF,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          &invalidCnpj,
				UserType:      1,
				ExpectedError: valueobject.ErrInvalidCNPJ,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          &validCnpj,
				UserType:      1,
				ExpectedError: ErrInvalidNaturalPerson,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      10,
				ExpectedError: valueobject.ErrInvalidUserType,
			},
			{
				Name:          "Paulo Silva",
				Email:         "paulo@gmail.com",
				Password:      "111A454aa",
				CPF:           "085.855.458-50",
				CNPJ:          nil,
				UserType:      2,
				ExpectedError: ErrInvalidLegalPerson,
			},
		}
		for _, item := range testCases {
			_, err := CreateUserFactory(nil, item.Name, item.Email, item.Password, item.CPF, item.CNPJ, item.UserType)
			if err != nil && !errors.Is(err, item.ExpectedError) {
				t.Errorf("Expected %f but got %f", item.ExpectedError, err)
			}
		}
	})
}
