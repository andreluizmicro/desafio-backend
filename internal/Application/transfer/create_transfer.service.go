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
	ErrUpdateBalance           = errors.New("error when try to update balance")
)

type CreateTransferService struct {
	accountRepository    contract.AccountRepositoryInterface
	transferRepository   contract.TransferRepositoryInterface
	authorizationGateway gateway.AuthorizationGateway
	notificationGateway  gateway.NotificationGatewayInterface
}

func NewCreateTransferService(
	accountRepository contract.AccountRepositoryInterface,
	transferRepository contract.TransferRepositoryInterface,
	authorizationGateway gateway.AuthorizationGateway,
	notificationGateway gateway.NotificationGatewayInterface,
) *CreateTransferService {
	return &CreateTransferService{
		accountRepository:    accountRepository,
		transferRepository:   transferRepository,
		authorizationGateway: authorizationGateway,
		notificationGateway:  notificationGateway,
	}
}

func (s *CreateTransferService) Execute(input CreateTransferInputDTO) (*CreateTransferOutputDTO, error) {
	accounts, err := s.getAccountsToTransaction(&input.Payer, &input.Payee)
	if err != nil {
		return nil, err
	}

	transferId, err := s.makeTransfer(input.Value, accounts["payer"], accounts["payee"])
	if err != nil {
		return nil, err
	}

	return &CreateTransferOutputDTO{
		ID: *transferId,
	}, nil
}

func (s *CreateTransferService) getAccountsToTransaction(payerId, payeeId *int64) (map[string]*entity.Account, error) {
	payer, err := s.accountRepository.FIndById(payerId)
	if err != nil {
		return nil, err
	}
	payee, err := s.accountRepository.FIndById(payeeId)
	if err != nil {
		return nil, err
	}
	accounts := map[string]*entity.Account{
		"payer": payer,
		"payee": payee,
	}
	return accounts, nil
}

func (s *CreateTransferService) makeTransfer(value float64, payer *entity.Account, payee *entity.Account) (*int64, error) {
	if !s.isAuthorized() {
		return nil, ErrUnauthorizedTransaction
	}
	transfer, err := entity.NewTransfer(nil, value, payer, payee)
	if err != nil {
		return nil, err
	}

	id, err := s.transferRepository.Create(transfer)
	if err != nil {
		return nil, err
	}

	err = s.UpdateUsersBalance(payer, payee)
	if err != nil {
		return nil, ErrUpdateBalance
	}

	if !s.notifyUsers() {
		return nil, ErrNotifyTransaction
	}
	return id, nil
}

func (s *CreateTransferService) isAuthorized() bool {
	return s.authorizationGateway.AuthorizeTransfer()
}

func (s *CreateTransferService) notifyUsers() bool {
	return s.notificationGateway.Notify()
}

func (s *CreateTransferService) UpdateUsersBalance(payer *entity.Account, payee *entity.Account) error {
	err := s.accountRepository.UpdateUserBalance(payer)
	if err != nil {
		return err
	}
	return s.accountRepository.UpdateUserBalance(payee)
}
