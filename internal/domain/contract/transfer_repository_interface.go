package contract

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	valueobject "github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

type TransferRepositoryInterface interface {
	Create(transfer entity.Transfer) (*valueobject.ID, error)
}
