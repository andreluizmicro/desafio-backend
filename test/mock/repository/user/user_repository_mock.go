package user

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Create(user *entity.User) (*int64, error) {
	args := m.Called(user)
	return args.Get(0).(*int64), args.Error(1)
}

func (m *RepositoryMock) FindByID(id *int64) (*entity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.User), args.Error(1)
}
