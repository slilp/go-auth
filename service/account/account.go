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
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
}

type SignInDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//go:generate mockgen -destination=../../mocks/mock_service/mock_account_service.go -package=mocks "github.com/slilp/go-auth/service/account" AccountService
type AccountService interface {
	CreateAccount(CreateAccountDto) (*AccountInfo, error)
	UpdateAccount(UpdateAccountDto) (AccountInfo, error)
	DeleteAccount(username string) error
	GetAccount(username string) (*AccountInfo, error)
	SignIn(SignInDto) (*AccountInfo, error)
}
