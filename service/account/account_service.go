package service

import (
	"github.com/jinzhu/copier"
	repository "github.com/slilp/go-auth/repository/account"
	utils "github.com/slilp/go-auth/utils"
)

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo}
}

func (s accountService) CreateAccount(req CreateAccountDto) (*AccountInfo, error) {
	if _, err := s.accountRepo.GetByUsername(req.Username); err == nil {
		return nil, utils.BadRequest("Duplicate account")
	}
	req.Password = utils.GenerateEncryptedPassword(req.Password)
	var createEntity repository.AccountEntity

	copier.Copy(&createEntity, &req)

	var srvRes AccountInfo
	repoRes, err := s.accountRepo.Create(createEntity)
	if err != nil {
		return nil, err
	}

	copier.Copy(&srvRes, repoRes)
	return &srvRes, err
}

func (s accountService) UpdateAccount(UpdateAccountDto) (AccountInfo, error) {
	return AccountInfo{}, nil
}

func (s accountService) DeleteAccount(username string) error {
	return nil
}

func (s accountService) GetAccount(username string) (*AccountInfo, error) {
	repoRes, err := s.accountRepo.GetByUsername(username)
	if err != nil {
		return nil, utils.NotFound("Not found account")
	}
	var srvRes AccountInfo
	copier.Copy(&srvRes, repoRes)
	return &srvRes, nil
}

func (s accountService) SignIn(req SignInDto) (*AccountInfo, error) {
	repoRes, err := s.accountRepo.GetByUsername(req.Username)
	if err != nil {
		return nil, utils.BadRequest("Invalid username or password")
	}

	if checkPassword := utils.ValidateEncryptedPassword(req.Password, repoRes.Password); !checkPassword {
		return nil, utils.BadRequest("Invalid username or password")
	}

	var srvRes AccountInfo
	copier.Copy(&srvRes, repoRes)
	return &srvRes, nil
}
