package account

import (
	"crud/entity"
	"crud/middleware"
	"crud/repository"
	"time"
)

type UseCaseAccount interface {
	CreateAccount(account AccountParam) (entity.Account, error)
	GetAccountById(id uint) (entity.Account, error)
	UpdateAccount(account AccountParam, id uint) (any, error)
	DeleteAccount(username string) (any, error)
	GetAccountByUsername(username string) (entity.Account, error)
	VerifyAccount(account AccountParam, id uint) (entity.Account, error)
	ActivateAccount(account AccountParam, id uint) (entity.Account, error)
}

type VerifiedStatus bool

const (
	True  VerifiedStatus = true
	False VerifiedStatus = false
)

type useCaseAccount struct {
	accountRepo repository.AccountInterfaceRepo
}

func (uc useCaseAccount) CreateAccount(account AccountParam) (entity.Account, error) {
	var newAccount *entity.Account

	// Hash the account password
	hashedPassword, err := middleware.HashPassword(account.Password)
	if err != nil {
		return entity.Account{}, err
	}

	newAccount = &entity.Account{
		Username: account.Username,
		Password: hashedPassword,
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err = uc.accountRepo.CreateAccount(newAccount)
	if err != nil {
		return *newAccount, err
	}
	return *newAccount, nil
}

func (uc useCaseAccount) GetAccountById(id uint) (entity.Account, error) {
	var account entity.Account
	account, err := uc.accountRepo.GetAccountsById(id)
	return account, err
}

func (uc useCaseAccount) UpdateAccount(account AccountParam, id uint) (any, error) {
	var editAccount *entity.Account
	editAccount = &entity.Account{
		ID:       account.ID,
		Username: account.Username,
		Password: account.Password,
		RoleID:   account.RoleID,
		Verified: account.Verified, // Use the specific enumeration value for true
		Active:   account.Active,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err := uc.accountRepo.UpdateAccount(editAccount)
	if err != nil {
		return *editAccount, err
	}
	return *editAccount, nil
}

func (uc useCaseAccount) DeleteAccount(email string) (any, error) {
	_, err := uc.accountRepo.DeleteAccount(email)
	return nil, err
}

func (uc useCaseAccount) GetAccountByUsername(username string) (entity.Account, error) {
	var account entity.Account
	account, err := uc.accountRepo.GetAccountsByUsername(username)
	return account, err
}

func (uc useCaseAccount) VerifyAccount(account AccountParam, id uint) (entity.Account, error) {
	var verifyAccount *entity.Account
	verifyAccount = &entity.Account{
		ID:       account.ID,
		Username: account.Username,
		Password: account.Password,
		RoleID:   account.RoleID,
		Verified: "true", // Use the specific enumeration value for true
		Active:   account.Active,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err := uc.accountRepo.VerifyAccount(verifyAccount)
	if err != nil {
		return *verifyAccount, err
	}
	return *verifyAccount, nil
}

func (uc useCaseAccount) ActivateAccount(account AccountParam, id uint) (entity.Account, error) {
	var activateAccount *entity.Account
	activateAccount = &entity.Account{
		ID:       account.ID,
		Username: account.Username,
		Password: account.Password,
		RoleID:   account.RoleID,
		Verified: account.Verified,
		Active:   "true",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	_, err := uc.accountRepo.ActivateAccount(activateAccount)
	if err != nil {
		return *activateAccount, err
	}
	return *activateAccount, nil
}
