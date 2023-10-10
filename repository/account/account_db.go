package repository

import (
	"gorm.io/gorm"
)

type accountRepositoryDB struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.AutoMigrate(&AccountEntity{})
	return accountRepositoryDB{db}
}

func (r accountRepositoryDB) Create(acc AccountEntity) (AccountEntity, error) {
	tx := r.db.Create(&acc)
	return acc, tx.Error
}

func (r accountRepositoryDB) GetByUsername(username string) (*AccountEntity, error) {
	var account AccountEntity
	tx := r.db.First(&account, "username = ?", username)
	return &account, tx.Error
}

func (r accountRepositoryDB) Delete(accId uint) error {
	return r.db.Delete(&AccountEntity{}, accId).Error
}

func (r accountRepositoryDB) Update(acc AccountEntity) error {
	return r.db.Model(&acc).Updates(acc).Error
}
