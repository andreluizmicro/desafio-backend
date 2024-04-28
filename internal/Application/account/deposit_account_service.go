package account

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/contract"
)

type DepositAccountService struct {
	accountRepository contract.AccountRepositoryInterface
}

func NewDepositAccountService(accountRepository contract.AccountRepositoryInterface) *DepositAccountService {
	return &DepositAccountService{
		accountRepository: accountRepository,
	}
}

func (s *DepositAccountService) Execute(input DepositAccountInputDto) (*DepositAccountOutputDto, error) {
	account, err := s.accountRepository.FIndById(&input.UserId)
	if err != nil || account == nil {
		return &DepositAccountOutputDto{Success: false}, err
	}
	err = account.CreditAccount(input.Value)
	if err != nil {
		return &DepositAccountOutputDto{Success: false}, err
	}
	err = s.accountRepository.UpdateUserBalance(account)
	if err != nil {
		return &DepositAccountOutputDto{Success: false}, err
	}

	return &DepositAccountOutputDto{Success: true}, nil
}
