package transfer

import (
	"errors"
	"github.com/andreluizmicro/desafio-backend/internal/domain/contract"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/gateway"
)

var (
	ErrUnauthorizedTransaction = errors.New("unauthorized transaction")
	ErrNotifyTransaction       = errors.New("error when notifying users transaction")
)

type CreateTransferService struct {
	accountRepository    contract.AccountRepositoryInterface
	transferRepository   contract.TransferRepositoryInterface
	authorizationGateway gateway.AuthorizationGateway
	NotificationGateway  gateway.NotificationGatewayInterface
}

func NewCreateTransferService(
	accountRepository contract.AccountRepositoryInterface,
	transferRepository contract.TransferRepositoryInterface,
	authorizationGateway gateway.AuthorizationGateway,
	NotificationGateway gateway.NotificationGatewayInterface,
) *CreateTransferService {
	return &CreateTransferService{
		accountRepository:    accountRepository,
		transferRepository:   transferRepository,
		authorizationGateway: authorizationGateway,
		NotificationGateway:  NotificationGateway,
	}
}

func (s *CreateTransferService) Execute(input CreateTransferInputDTO) (*CreateTransferOutputDTO, error) {
	payer, err := s.accountRepository.FIndById(&input.Payer)
	if err != nil {
		return nil, err
	}
	payee, err := s.accountRepository.FIndById(&input.Payee)
	if err != nil {
		return nil, err
	}

	if !s.isAuthorized() {
		return nil, ErrUnauthorizedTransaction
	}
	transfer, err := entity.NewTransfer(nil, input.Value, payer, payee)
	if err != nil {
		return nil, err
	}

	id, err := s.transferRepository.Create(transfer)
	if err != nil {
		return nil, err
	}

	if !s.notifyUsers() {
		return nil, ErrNotifyTransaction
	}

	return &CreateTransferOutputDTO{
		ID: *id,
	}, nil
}

func (s *CreateTransferService) isAuthorized() bool {
	return s.authorizationGateway.AuthorizeTransfer()
}

func (s *CreateTransferService) notifyUsers() bool {
	return s.NotificationGateway.Notify()
}
