package account

type CreateAccountInputDto struct {
	UserId int64 `json:"user_id"`
}

type CreateAccountOutputDto struct {
	Id int64 `json:"id"`
}

type DepositAccountInputDto struct {
	UserId int64   `uri:"user_id"`
	Value  float64 `json:"value"`
}

type DepositAccountOutputDto struct {
	Success bool `json:"success"`
}
