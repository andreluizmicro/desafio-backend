package account

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/exception"
)

type Repository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(account *entity.Account) (*int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, user_id, balance) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(
		account.ID(),
		*account.User().ID,
		account.Balance(),
	)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (r *Repository) FIndById(id *int64) (*entity.Account, error) {
	var account entity.Account
	stmt, err := r.db.Prepare(`
			SELECT * FROM accounts 
			INNER JOIN users ON (accounts.user_id = users.id)         
         	WHERE id = ? AND active = ?
		`,
	)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id, 1).Scan(
		account.ID(),
		account.User,
		account.Balance,
	)
	if err != nil {
		return nil, exception.ErrUserNotFound
	}
	return &account, nil
}
