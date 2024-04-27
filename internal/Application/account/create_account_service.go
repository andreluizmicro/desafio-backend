package account

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/contract"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
)

type CreateAccountService struct {
	accountRepository contract.AccountRepositoryInterface
	userRepository    contract.UserRepositoryInterface
}

func NewCreateAccountService(
	accountRepository contract.AccountRepositoryInterface,
	userRepository contract.UserRepositoryInterface,
) *CreateAccountService {
	return &CreateAccountService{
		accountRepository: accountRepository,
		userRepository:    userRepository,
	}
}

func (s *CreateAccountService) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {

	user, err := s.userRepository.FindByID(&input.UserId)
	if err != nil {
		return nil, err
	}

	account, err := entity.NewAccount(nil, user, 0)
	if err != nil {
		return nil, err
	}

	id, err := s.accountRepository.Create(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDto{
		Id: *id,
	}, nil
}
