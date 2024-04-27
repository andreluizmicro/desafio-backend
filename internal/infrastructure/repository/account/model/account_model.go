package model

type AccountModel struct {
	AccountID    int64   `json:"id"`
	Balance      float64 `json:"balance"`
	UserID       int64   `json:"user_id"`
	UserName     string  `json:"name"`
	UserPassword string  `json:"password"`
	UserEmail    string  `json:"email"`
	UserCPF      string  `json:"cpf"`
	UserCNPJ     string  `json:"cnpj"`
	UserTypeID   int64   `json:"user_type"`
}
