package contract

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	valueobject "github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

type AccountRepositoryInterface interface {
	Create(account *entity.Account) (*valueobject.ID, error)
	FIndById(id valueobject.ID) (*entity.Account, error)
}
