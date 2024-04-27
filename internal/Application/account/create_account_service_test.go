package account

//func TestCreateAccountService(t *testing.T) {
//	accountRepositoryMock := &accountRepository.RepositoryMock{}
//	userRepositoryMock := &userRepository.RepositoryMock{}
//
//	createAccountService := NewCreateAccountService(accountRepositoryMock, userRepositoryMock)
//	user, _ := entity.CreateUserFactory(
//		nil,
//		"André Silva",
//		"andre@gmail.com",
//		"12345AAaa@",
//		"088.445.458-65",
//		nil,
//		1,
//	)
//	id := valueobject.NewID()
//
//	accountRepositoryMock.On("Create", mock.Anything).Return(id, nil).Once()
//	userRepositoryMock.On("FindById", mock.Anything).Return(user, nil).Once()
//
//	_, err := createAccountService.Create(CreateAccountInputDto{
//		UserId: id.Value,
//	})
//	fmt.Println(err)
//	//assert.Nil(t, err)
//	//assert.NotNil(t, output)
//	//assert.Equal(t, id, output.Id)
//}
