package repository

import "gorm.io/gorm"

type AccountEntity struct {
	gorm.Model
	Username  string `gorm:"unique_index;not null;"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	Avatar    string
	Role      string `gorm:"default:'USER';not null;"`
}

func (AccountEntity) TableName() string {
	return "account"
}

//go:generate mockgen -destination=../../mocks/mock_repository/mock_account_repository.go -package=mocks "github.com/slilp/go-auth/repository/account" AccountRepository
type AccountRepository interface {
	GetByUsername(username string) (*AccountEntity, error)
	Create(AccountEntity) (AccountEntity, error)
	Delete(id string) error
	Update(AccountEntity) (AccountEntity, error)
}
