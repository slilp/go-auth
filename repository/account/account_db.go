package repository

import "gorm.io/gorm"

type accountRepositoryDB struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return accountRepositoryDB{db}
}

func (r accountRepositoryDB) Create(acc AccountEntity) (AccountEntity, error) {
	return AccountEntity{}, nil
}

func (r accountRepositoryDB) GetByUsername(username string) (AccountEntity, error) {
	return AccountEntity{}, nil
}

func (r accountRepositoryDB) Delete(accId string) error {
	return nil
}

func (r accountRepositoryDB) Update(acc AccountEntity) (AccountEntity, error) {
	return AccountEntity{}, nil
}
