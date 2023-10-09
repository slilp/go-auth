package service

import repository "github.com/slilp/go-auth/repository/account"

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return accountService{accountRepo}
}

func (s accountService) CreateAccount(CreateAccountDto) (AccountInfo, error) {
	return AccountInfo{}, nil
}
func (s accountService) UpdateAccount(UpdateAccountDto) (AccountInfo, error) {
	return AccountInfo{}, nil

}
func (s accountService) DeleteAccount(username string) error {
	return nil
}
func (s accountService) GetAccount(username string) (AccountInfo, error) {
	return AccountInfo{}, nil

}
