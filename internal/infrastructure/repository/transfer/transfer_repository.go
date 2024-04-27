package transfer

import (
	"database/sql"

	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
)

type Repository struct {
	db *sql.DB
}

func NewTransferRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(transfer *entity.Transfer) (*int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO transfers (payer_is, payee_id, value) VALUES (?, ?, ?)")

	if err != nil {
		return new(int64), err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		transfer.Payer(),
		transfer.Payee(),
		transfer.Value(),
	)
	if err != nil {
		return new(int64), err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return new(int64), err
	}
	return &id, err
}
