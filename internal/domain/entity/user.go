package entity

import (
	"errors"
	"time"

	"github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
)

const (
	defaultUser    = 1
	shopkeeperUser = 2
)

var (
	ErrInvalidNaturalPerson = errors.New("natural person does not have a CNPJ")
	ErrInvalidLegalPerson   = errors.New("legal entity must have a CNPJ")
)

type User struct {
	ID        *value_object.ID       `json:"id"`
	Name      *value_object.Name     `json:"name"`
	Email     *value_object.Email    `json:"email"`
	Password  *value_object.Password `json:"-"`
	CPF       *value_object.CPF      `json:"cpf"`
	CNPJ      *value_object.CNPJ     `json:"cnpj"`
	UserType  *value_object.UserType `json:"user_type"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Deleted   bool                   `json:"deleted"`
}

func newUser(
	id *value_object.ID,
	name *value_object.Name,
	email *value_object.Email,
	password *value_object.Password,
	cpf *value_object.CPF,
	cnpj *value_object.CNPJ,
	userType *value_object.UserType,
) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
		CPF:      cpf,
		CNPJ:     cnpj,
		UserType: userType,
	}
	user.ID = id
	err := user.validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUserFactory(id *string, name, email, password, cpf string, cnpj *string, userTypeId int) (*User, error) {
	var userId *value_object.ID
	if id == nil {
		userId = value_object.NewID()
	}
	userName, err := value_object.NewName(name)
	if err != nil {
		return nil, err
	}
	userEmail, err := value_object.NewEmail(email)
	if err != nil {
		return nil, err
	}
	userPassword, err := value_object.NewPassword(password)
	if err != nil {
		return nil, err
	}
	userCpf, err := value_object.NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	var userCnpj *value_object.CNPJ
	if cnpj != nil {
		cnpj, err := value_object.NewCNPJ(*cnpj)
		if err != nil {
			return nil, err
		}
		userCnpj = cnpj
	}
	userType := value_object.NewUserType(userTypeId)

	return newUser(
		userId,
		userName,
		userEmail,
		userPassword,
		userCpf,
		userCnpj,
		userType,
	)
}

func (u *User) validate() error {
	if u.UserType.Value == defaultUser && u.CNPJ != nil {
		return ErrInvalidNaturalPerson
	}
	if u.UserType.Value == shopkeeperUser && u.CNPJ == nil {
		return ErrInvalidLegalPerson
	}

	return nil
}
