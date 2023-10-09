package service

type CreateAccountDto struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
}

type UpdateAccountDto struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
}

type AccountInfo struct {
	Username  int     `json:"username"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Avatar    float64 `json:"avatar"`
	Role      int     `json:"role"`
}

//go:generate mockgen -destination=../../mock/mock_service/mock_account_service.go "github.com/slilp/go-auth" AccountService
type AccountService interface {
	CreateAccount(CreateAccountDto) (AccountInfo, error)
	UpdateAccount(UpdateAccountDto) (AccountInfo, error)
	DeleteAccount(username string) error
	GetAccount(username string) (AccountInfo, error)
}
