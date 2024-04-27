package repository_mock

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *entity.User) (*value_object.ID, error) {
	args := m.Called(user)
	return args.Get(0).(*value_object.ID), args.Error(1)
}
