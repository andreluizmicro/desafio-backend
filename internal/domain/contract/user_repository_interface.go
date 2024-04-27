package contract

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

type UserRepositoryInterface interface {
	Create(user *entity.User) (*value_object.ID, error)
}
