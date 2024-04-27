package transfer

type CreateTransferInputDTO struct {
	Value float64 `json:"value"`
	Payer int64   `json:"payer"`
	Payee int64   `json:"payee"`
}

type CreateTransferOutputDTO struct {
	ID int64 `json:"id"`
}
