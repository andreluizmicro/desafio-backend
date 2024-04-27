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

func (r *Repository) Create(user *entity.User) (*value_object.ID, error) {
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

func (r *Repository) FindByID(id value_object.ID) (*entity.User, error) {
	var userModel model.UserModel
	stmt, err := r.db.Prepare("SELECT id, name, email, password, cpf, cnpj, user_type_id password FROM users WHERE id = ? AND deleted = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id.Value, 0).Scan(
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
