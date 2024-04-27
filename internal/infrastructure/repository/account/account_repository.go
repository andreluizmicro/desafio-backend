package account

import (
	"database/sql"
	"fmt"

	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/exception"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

type Repository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(account *entity.Account) (*value_object.ID, error) {
	stmt, err := r.db.Prepare("INSERT INTO accounts (id, user_id, balance) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}

	fmt.Println("@!DSDADSADSADSAHDSAHDJSAHDJKSAHJKDASDJSA")

	defer stmt.Close()
	_, err = stmt.Exec(
		account.ID().Value,
		"9694b815-3e32-4d8d-b621-e79a4b890a49",
		account.Balance(),
	)
	if err != nil {
		return nil, err
	}
	return account.ID(), nil
}

func (r *Repository) FIndById(id value_object.ID) (*entity.Account, error) {
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
		&account.ID().Value,
		account.User,
		account.Balance,
	)
	if err != nil {
		return nil, exception.ErrUserNotFound
	}

	return &account, nil
}
