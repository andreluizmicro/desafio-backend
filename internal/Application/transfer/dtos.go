package transfer

type CreateTransferInputDTO struct {
	Value float64 `json:"value"`
	Payer string  `json:"payer"`
	Payee string  `json:"payee"`
}

type CreateTransferOutputDTO struct {
	ID string `json:"id"`
}
