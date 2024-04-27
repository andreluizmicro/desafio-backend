package account

type CreateAccountInputDto struct {
	UserId int64 `json:"user_id"`
}

type CreateAccountOutputDto struct {
	Id int64 `json:"id"`
}
