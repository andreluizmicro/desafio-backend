package user

import (
	"database/sql"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/exception"
	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
	"github.com/andreluizmicro/desafio-backend/internal/infrastructure/repository/user/model"
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(user *entity.User) (*int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email, cpf, cnpj, password, user_type_id) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	var cnpj *value_object.CNPJ
	if user.CNPJ == nil {
		cnpj = nil
	}

	defer stmt.Close()
	result, err := stmt.Exec(
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

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &id, nil
}

func (r *Repository) FindByID(id *int64) (*entity.User, error) {
	var userModel model.UserModel
	stmt, err := r.db.Prepare("SELECT id, name, email, password, cpf, cnpj, user_type_id password FROM users WHERE id = ? AND deleted = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(*id, 0).Scan(
		&userModel.ID,
		&userModel.Name,
		&userModel.Email,
		&userModel.Password,
		&userModel.CPF,
		&userModel.CNPJ,
		&userModel.UserTypeId,
	)
	if err != nil {
		return nil, exception.ErrUserNotFound
	}
	return modelToEntity(userModel)
}

func modelToEntity(user model.UserModel) (*entity.User, error) {
	return entity.CreateUserFactory(
		user.ID,
		user.Name,
		user.Email,
		user.Password,
		user.CPF,
		user.CNPJ,
		user.UserTypeId,
	)
}
