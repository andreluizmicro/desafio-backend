package user

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	repositorymock "github.com/andreluizmicro/desafio-backend/test/mock/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestCreateUser(t *testing.T) {
	repositoryMock := &repositorymock.UserRepositoryMock{}
	createUserService := NewCreateUserService(repositoryMock)

	t.Run("test should create user", func(t *testing.T) {
		input := CreateUserInputDto{
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "Password123A@s",
			CPF:      "088.988.888-52",
			CNPJ:     nil,
			UserType: 1,
		}

		repositoryMock.On("Create", mock.Anything).Return(value_object.NewID(), nil).Once()
		output, err := createUserService.Execute(input)
		assert.NoError(t, err)
		assert.NotNil(t, output)
		assert.NotEmpty(t, output.Id)
	})

	t.Run("test should return invalid legal person", func(t *testing.T) {
		var id *value_object.ID
		repositoryMock.On("Create", mock.Anything).Return(id, entity.ErrInvalidLegalPerson).Once()
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
