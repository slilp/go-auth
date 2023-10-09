package repository

import "gorm.io/gorm"

type AccountEntity struct {
	gorm.Model
	Username  string `gorm:"unique_index;"`
	Password  string `gorm:"not null"`
	FirstName string
	LastName  string
	Avatar    string
	Role      string `gorm:"default:'USER'"`
}

func (AccountEntity) TableName() string {
	return "account"
}

//go:generate mockgen -destination=../../mock/mock_repository/mock_account_repository.go "github.com/slilp/go-auth" AccountRepository
type AccountRepository interface {
	GetByUsername(username string) (AccountEntity, error)
	Create(AccountEntity) (AccountEntity, error)
	Delete(id string) error
	Update(AccountEntity) (AccountEntity, error)
}
