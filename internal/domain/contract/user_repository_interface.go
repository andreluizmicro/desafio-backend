package contract

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
)

type UserRepositoryInterface interface {
	Create(user *entity.User) (*int64, error)
	FindByID(id *int64) (*entity.User, error)
}
