package account

import (
	"crud/entity"
	"crud/repository/mocks"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestCreateAccount(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	accountParam := AccountParam{
		Username: "testaccount",
		Password: "password",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(accountParam.Password), bcrypt.DefaultCost)
	expectedAccount := entity.Account{
		Username: accountParam.Username,
		Password: string(hashedPassword),
		RoleID:   2,
		Verified: "false",
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the CreateAccount method
	accountRepo.On("CreateAccount", mock.AnythingOfType("*entity.Account")).Return(&expectedAccount, nil)

	createdAccount, err := accountUseCase.CreateAccount(accountParam)

	// Compare the hashed password with the provided password
	err = bcrypt.CompareHashAndPassword([]byte(createdAccount.Password), []byte(accountParam.Password))
	assert.NoError(t, err)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount.Username, createdAccount.Username)
	assert.Equal(t, expectedAccount.RoleID, createdAccount.RoleID)

}

func TestDeleteAccount(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	email := "test@example.com"

	// Mock the DeleteAccount method
	accountRepo.On("DeleteAccount", email).Return(nil, nil)

	_, err := accountUseCase.DeleteAccount(email)

	assert.NoError(t, err)
}

func TestUpdateAccount(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	accountParam := AccountParam{
		ID:       1,
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "true", // Use the specific enumeration value for true
		Active:   "false",
	}

	editAccount := &entity.Account{
		ID:       accountParam.ID,
		Username: accountParam.Username,
		Password: accountParam.Password,
		RoleID:   accountParam.RoleID,
		Verified: accountParam.Verified,
		Active:   accountParam.Active,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the UpdateAccount method
	accountRepo.On("UpdateAccount", editAccount).Return(editAccount, nil)

	updatedAccount, err := accountUseCase.UpdateAccount(accountParam, accountParam.ID)

	assert.NoError(t, err)
	assert.Equal(t, *editAccount, updatedAccount)
	//fmt.Println(editAccount.Password)
}

func TestGetAccountById(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	accountID := uint(1)

	expectedAccount := entity.Account{
		ID:       accountID,
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "true", // Use the specific enumeration value for true
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the GetAccountsById method
	accountRepo.On("GetAccountsById", accountID).Return(expectedAccount, nil)

	retrievedAccount, err := accountUseCase.GetAccountById(accountID)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, retrievedAccount)

	// Print the retrieved account
	fmt.Println("Retrieved Account:", retrievedAccount)
}

func TestGetAccountByUsername(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	username := "testuser"

	expectedAccount := entity.Account{
		ID:       1,
		Username: username,
		Password: "testpassword",
		RoleID:   2,
		Verified: "true", // Use the specific enumeration value for true
		Active:   "false",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the GetAccountsByUsername method
	accountRepo.On("GetAccountsByUsername", username).Return(expectedAccount, nil)

	retrievedAccount, err := accountUseCase.GetAccountByUsername(username)

	assert.NoError(t, err)
	assert.Equal(t, expectedAccount, retrievedAccount)

	// Print the retrieved account
	fmt.Println("Retrieved Account:", retrievedAccount)
}

func TestVerifyAccount(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	accountParam := AccountParam{
		ID:       1,
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "true", // Use the specific enumeration value for true
		Active:   "false",
	}

	verifyAccount := &entity.Account{
		ID:       accountParam.ID,
		Username: accountParam.Username,
		Password: accountParam.Password,
		RoleID:   accountParam.RoleID,
		Verified: accountParam.Verified,
		Active:   accountParam.Active,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the VerifyAccount method
	accountRepo.On("VerifyAccount", verifyAccount).Return(verifyAccount, nil)

	verifiedAccount, err := accountUseCase.VerifyAccount(accountParam, accountParam.ID)

	assert.NoError(t, err)
	assert.Equal(t, *verifyAccount, verifiedAccount)

	// Print the verified account
	fmt.Println("Verified Account:", verifiedAccount)
}

func TestActivateAccount(t *testing.T) {
	accountRepo := new(mocks.AccountInterfaceRepo)
	accountUseCase := useCaseAccount{
		accountRepo: accountRepo,
	}

	accountParam := AccountParam{
		ID:       1,
		Username: "testuser",
		Password: "testpassword",
		RoleID:   2,
		Verified: "true", // Use the specific enumeration value for true
		Active:   "false",
	}

	activateAccount := &entity.Account{
		ID:       accountParam.ID,
		Username: accountParam.Username,
		Password: accountParam.Password,
		RoleID:   accountParam.RoleID,
		Verified: accountParam.Verified,
		Active:   "true",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the ActivateAccount method
	accountRepo.On("ActivateAccount", activateAccount).Return(activateAccount, nil)

	activatedAccount, err := accountUseCase.ActivateAccount(accountParam, accountParam.ID)

	assert.NoError(t, err)
	assert.Equal(t, *activateAccount, activatedAccount)

	// Print the activated account
	fmt.Println("Activated Account:", activatedAccount)
}
