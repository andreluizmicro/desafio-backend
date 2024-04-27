package account

type CreateAccountInputDto struct {
	UserId string `json:"user_id"`
}

type CreateAccountOutputDto struct {
	Id string `json:"id"`
}
