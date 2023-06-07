package repository

import (
	"crud/entity"
	"gorm.io/gorm"
)

type Account struct {
	db *gorm.DB
}

func NewAccount(dbCrud *gorm.DB) Account {
	return Account{
		db: dbCrud,
	}

}

type AccountInterfaceRepo interface {
	CreateAccount(account *entity.Account) (*entity.Account, error)
	GetAccountsById(id uint) (entity.Account, error)
	UpdateAccount(account *entity.Account) (any, error)
	DeleteAccount(username string) (any, error)
	GetAccountsByUsername(username string) (entity.Account, error)
	VerifyAccount(account *entity.Account) (any, error)
	ActivateAccount(account *entity.Account) (any, error)
}

func (repo Account) CreateAccount(account *entity.Account) (*entity.Account, error) {
	err := repo.db.Model(&entity.Account{}).Create(account).Error
	return account, err
}

func (repo Account) GetAccountsById(id uint) (entity.Account, error) {
	var account entity.Account
	repo.db.First(&account, "id = ? ", id)
	return account, nil
}

// UpdateUser multiple fields
func (repo Account) UpdateAccount(account *entity.Account) (any, error) {
	err := repo.db.Model(&entity.Account{}).Preload("Role").
		Where("id = ?", account.ID).
		Updates(&account).Error
	return nil, err
}

// DeleteUser by Id and email
func (repo Account) DeleteAccount(username string) (any, error) {
	err := repo.db.Model(&entity.Account{}).
		Where("username = ?", username).
		Delete(&entity.Account{}).
		Error
	return nil, err
}

func (repo Account) GetAccountsByUsername(username string) (entity.Account, error) {
	var account entity.Account
	repo.db.First(&account, "username = ? ", username)
	return account, nil
}

func (repo Account) VerifyAccount(account *entity.Account) (any, error) {
	err := repo.db.Model(&entity.Account{}).
		Where("id = ?", account.ID).
		Update("verified", false).Error // Update the 'verified' column to false
	return nil, err
}

func (repo Account) ActivateAccount(account *entity.Account) (any, error) {
	err := repo.db.Model(&entity.Account{}).
		Where("id = ?", account.ID).
		Update("verified", false).Error // Update the 'verified' column to false
	return nil, err
}
