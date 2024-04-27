package user

type CreateUserInputDto struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	CPF      string  `json:"cpf"`
	CNPJ     *string `json:"cnpj"`
	Password string  `json:"password"`
	UserType int     `json:"user_type_id"`
}

type CreateUserOutputDto struct {
	Id *int64 `json:"id"`
}
