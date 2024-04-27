package user

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	repositorymock "github.com/andreluizmicro/desafio-backend/test/mock/repository/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUserService(t *testing.T) {
	repositoryMock := &repositorymock.RepositoryMock{}
	createUserService := NewCreateUserService(repositoryMock)
	var id int64 = 10
	t.Run("test should create user", func(t *testing.T) {
		input := CreateUserInputDto{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Password123A@s",
			CPF:      "088.988.888-52",
			CNPJ:     nil,
			UserType: 1,
		}

		repositoryMock.On("Create", mock.Anything).Return(&id, nil).Once()
		output, err := createUserService.Execute(input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
	})

	t.Run("test should return invalid legal person", func(t *testing.T) {
		repositoryMock.On("Create", mock.Anything).Return(&id, entity.ErrInvalidLegalPerson).Once()
		output, err := createUserService.Execute(
			CreateUserInputDto{
				Name:     "John Doe",
				Email:    "john.doe@example.com",
				Password: "Password123A@s",
				CPF:      "088.988.888-52",
				CNPJ:     nil,
				UserType: 1,
			},
		)
		assert.Error(t, err)
		assert.Nil(t, output)
	})

	t.Run("test should return invalid email", func(t *testing.T) {
		var id *value_object.ID
		repositoryMock.On("Create", mock.Anything).Return(id, value_object.ErrInvalidEmail).Once()
		output, err := createUserService.Execute(
			CreateUserInputDto{
				Name:     "John Doe",
				Email:    "aba",
				Password: "Password123A@s",
				CPF:      "088.988.888-52",
				CNPJ:     nil,
				UserType: 1,
			},
		)
		assert.Error(t, err)
		assert.Nil(t, output)
	})
}
