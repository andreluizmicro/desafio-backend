package user

type CreateUserInputDto struct {
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	CPF      string  `json:"cpf"`
	CNPJ     *string `json:"cnpj"`
	Password string  `json:"password"`
	UserType int     `json:"user_type"`
}

type CreateUserOutputDto struct {
	Id string `json:"id"`
}
