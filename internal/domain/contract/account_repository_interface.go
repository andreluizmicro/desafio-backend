package contract

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
)

type AccountRepositoryInterface interface {
	Create(account *entity.Account) (*int64, error)
	FIndById(id *int64) (*entity.Account, error)
	ExistsById(id *int64) bool
}
