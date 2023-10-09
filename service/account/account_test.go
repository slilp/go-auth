package service_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	mocks "github.com/slilp/go-auth/mocks/mock_repository"
	repository "github.com/slilp/go-auth/repository/account"
	service "github.com/slilp/go-auth/service/account"
	utils "github.com/slilp/go-auth/utils"
	"github.com/stretchr/testify/assert"
)

func newFakeAccount() repository.AccountEntity {
	return repository.AccountEntity{
		Username:  "faleUsername",
		Password:  "fakePassword",
		FirstName: "fakeFirstName",
		LastName:  "fakeLastName",
		Avatar:    "fakeAvatar",
		Role:      "fakeRole",
	}
}

func TestGetAccount(t *testing.T) {

	t.Run("Not found error", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)

		mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
		mockAccountRepo.EXPECT().GetByUsername(gomock.Any().String()).Return(nil, assert.AnError)
		accountService := service.NewAccountService(mockAccountRepo)

		//Act
		_, err := accountService.GetAccount(gomock.Any().String())

		//Assert
		assert.Equal(t, utils.NotFound("Not found account"), err)
	})

	t.Run("Success", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)
		fakeAccount := newFakeAccount()

		mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
		mockAccountRepo.EXPECT().GetByUsername(gomock.Any().String()).Return(&fakeAccount, nil)
		accountService := service.NewAccountService(mockAccountRepo)

		//Act
		repoRes, err := accountService.GetAccount(gomock.Any().String())

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, &service.AccountInfo{
			Username:  fakeAccount.Username,
			FirstName: fakeAccount.FirstName,
			LastName:  fakeAccount.LastName,
			Avatar:    fakeAccount.Avatar,
			Role:      fakeAccount.Role,
		}, repoRes)

	})

}

func TestCreateAccount(t *testing.T) {

	t.Run("Duplicate username", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)
		fakeAccount := newFakeAccount()

		mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
		mockAccountRepo.EXPECT().GetByUsername(gomock.Any().String()).Return(&fakeAccount, nil)
		accountService := service.NewAccountService(mockAccountRepo)

		//Act
		_, err := accountService.CreateAccount(service.CreateAccountDto{Username: gomock.Any().String()})

		//Assert
		assert.Equal(t, utils.BadRequest("Duplicate account"), err)
	})

	t.Run("Save record error", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)
		fakeAccount := newFakeAccount()

		mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
		mockAccountRepo.EXPECT().GetByUsername(gomock.Any().String()).Return(nil, assert.AnError)
		mockAccountRepo.EXPECT().Create(gomock.Any()).Return(fakeAccount, assert.AnError)
		accountService := service.NewAccountService(mockAccountRepo)

		//Act
		_, err := accountService.CreateAccount(service.CreateAccountDto{Username: gomock.Any().String()})

		//Assert
		assert.Equal(t, assert.AnError, err)
	})

	t.Run("Success", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)
		fakeAccount := newFakeAccount()

		mockAccountRepo := mocks.NewMockAccountRepository(ctrl)
		mockAccountRepo.EXPECT().GetByUsername(gomock.Any().String()).Return(nil, assert.AnError)
		mockAccountRepo.EXPECT().Create(gomock.Any()).Return(fakeAccount, nil)
		accountService := service.NewAccountService(mockAccountRepo)

		//Act
		repoRes, err := accountService.CreateAccount(service.CreateAccountDto{Username: gomock.Any().String()})

		//Assert
		assert.NoError(t, err)
		assert.Equal(t, &service.AccountInfo{
			Username:  fakeAccount.Username,
			FirstName: fakeAccount.FirstName,
			LastName:  fakeAccount.LastName,
			Avatar:    fakeAccount.Avatar,
			Role:      fakeAccount.Role,
		}, repoRes)

	})

}
