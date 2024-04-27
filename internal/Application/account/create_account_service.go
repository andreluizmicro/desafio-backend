package account

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/contract"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	"github.com/andreluizmicro/desafio-backend/internal/domain/exception"
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
	if s.existsById(input.UserId) {
		return nil, exception.ErrAccountAlreadyExists
	}

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

func (s *CreateAccountService) existsById(id int64) bool {
	return s.accountRepository.ExistsById(&id)
}
