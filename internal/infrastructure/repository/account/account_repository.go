package account

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/repository/account/model"
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

func (r *Repository) ExistsById(id *int64) bool {
	var userId *int64
	stmt := `SELECT id FROM accounts WHERE user_id = ?`
	err := r.db.QueryRow(stmt, id).Scan(&userId)
	if err != nil {
		return false
	}
	return &id != nil
}

func (r *Repository) FIndById(id *int64) (*entity.Account, error) {
	stmt, err := r.db.Prepare(`
		SELECT 
		    a.id, 
            a.balance, 
            u.id as user_id,
            u.user_type_id,
            u.name,
            u.password,
            u.email,
            u.cpf,
            u.cnpj
		FROM accounts a 
		INNER JOIN users u ON (u.id = a.user_id) 
		WHERE u.id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var accountModel model.AccountModel

	_ = stmt.QueryRow(*id).Scan(
		&accountModel.AccountID,
		&accountModel.Balance,
		&accountModel.UserID,
		&accountModel.UserTypeID,
		&accountModel.UserName,
		&accountModel.UserPassword,
		&accountModel.UserEmail,
		&accountModel.UserCPF,
		&accountModel.UserCNPJ,
	)

	return modelToEntity(accountModel)
}

func modelToEntity(model model.AccountModel) (*entity.Account, error) {
	user, err := entity.CreateUserFactory(
		&model.UserID,
		model.UserName,
		model.UserEmail,
		model.UserPassword,
		model.UserCPF,
		nil,
		int(model.UserTypeID),
	)
	if err != nil {
		return nil, err
	}
	return entity.NewAccount(&model.AccountID, user, model.Balance)
}
