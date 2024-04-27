package account

import (
	"fmt"
	"github.com/andreluizmicro/desafio-backend/internal/domain/contract"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
	valueobject "github.com/andreluizmicro/desafio-backend/internal/domain/value_object"
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
	user, err := s.userRepository.FindByID(valueobject.ID{Value: input.UserId})
	if err != nil {
		return nil, err
	}

	fmt.Println("@!DSDADSADSADSAHDSAHDJSAHDJKSAHJKDASDJSA", user.ID)

	account, err := entity.NewAccount(nil, user, 0)
	if err != nil {
		return nil, err
	}

	_, err = s.accountRepository.Create(account)
	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDto{
		Id: account.ID().Value,
	}, nil
}
