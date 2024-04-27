package account

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) Create(account *entity.Account) (*int64, error) {
	args := m.Called(account)
	return args.Get(0).(*int64), args.Error(1)
}

func (m *RepositoryMock) FIndById(id value_object.ID) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}
