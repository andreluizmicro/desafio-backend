package user

import (
	"github.com/andreluizmicro/desafio-backend/internal/domain/contract"
	"github.com/andreluizmicro/desafio-backend/internal/domain/entity"
)

type CreateUserService struct {
	userRepository contract.UserRepositoryInterface
}

func NewCreateUserService(userRepository contract.UserRepositoryInterface) *CreateUserService {
	return &CreateUserService{userRepository: userRepository}
}

func (s *CreateUserService) Execute(input CreateUserInputDto) (*CreateUserOutputDto, error) {
	user, err := entity.CreateUserFactory(
		nil,
		input.Name,
		input.Email,
		input.Password,
		input.CPF,
		input.CNPJ,
		input.UserType,
	)
	if err != nil {
		return nil, err
	}

	_, err = s.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutputDto{
		Id: user.ID.Value,
	}, nil
}
