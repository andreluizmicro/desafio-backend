package model

type UserModel struct {
	ID         *string `json:"id"`
	Name       string  `json:"name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	CPF        string  `json:"cpf"`
	CNPJ       *string `json:"cnpj"`
	UserTypeId int     `json:"user_type_id"`
}
