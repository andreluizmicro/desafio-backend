package repository

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user *entity.User) (*value_object.ID, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (id, name, email, cpf, cnpj, password, user_type_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	var cnpj *value_object.CNPJ
	if user.CNPJ == nil {
		cnpj = nil
	}

	defer stmt.Close()
	_, err = stmt.Exec(
		user.ID.Value,
		user.Name.Value,
		user.Email.Value,
		user.CPF.Value,
		cnpj,
		user.Password.Value,
		user.UserType.Value,
	)
	if err != nil {
		return nil, err
	}
	return user.ID, nil
}
