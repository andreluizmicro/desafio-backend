package user

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Create(user *entity.User) (*value_object.ID, error) {
	args := m.Called(user)
	return args.Get(0).(*value_object.ID), args.Error(1)
}

func (m *RepositoryMock) FindByID(id value_object.ID) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}
