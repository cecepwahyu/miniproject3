package user

import (
	"crud/entity"
	"crud/repository/mocks"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	userRepo := new(mocks.UserInterfaceRepo)
	userUseCase := useCaseUser{
		userRepo: userRepo,
	}

	userParam := UserParam{
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
	}

	newUser := &entity.User{
		Name:     userParam.Name,
		Email:    userParam.Email,
		Password: userParam.Password,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the CreateUser method
	userRepo.On("CreateUser", newUser).Return(newUser, nil)

	createdUser, err := userUseCase.CreateUser(userParam)

	assert.NoError(t, err)
	assert.Equal(t, *newUser, createdUser)

	// Print the created user
	fmt.Println("Created User:", createdUser)
}

func TestGetUserById(t *testing.T) {
	userRepo := new(mocks.UserInterfaceRepo)
	userUseCase := useCaseUser{
		userRepo: userRepo,
	}

	userID := uint(1)

	expectedUser := entity.User{
		ID:       userID,
		Name:     "John Doe",
		Email:    "john.doe@example.com",
		Password: "password",
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the GetUsersById method
	userRepo.On("GetUsersById", userID).Return(expectedUser, nil)

	retrievedUser, err := userUseCase.GetUserById(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, retrievedUser)

	// Print the retrieved user
	fmt.Println("Retrieved User:", retrievedUser)
}

func TestUpdateUser(t *testing.T) {
	userRepo := new(mocks.UserInterfaceRepo)
	userUseCase := useCaseUser{
		userRepo: userRepo,
	}

	userID := uint(1)

	updatedUserParam := UserParam{
		Name:     "Updated John Doe",
		Email:    "updated.john.doe@example.com",
		Password: "updatedpassword",
	}

	updatedUser := &entity.User{
		ID:       userID,
		Name:     updatedUserParam.Name,
		Email:    updatedUserParam.Email,
		Password: updatedUserParam.Password,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}

	// Mock the UpdateUser method
	userRepo.On("UpdateUser", updatedUser).Return(updatedUser, nil)

	editedUser, err := userUseCase.UpdateUser(updatedUserParam, userID)

	assert.NoError(t, err)
	assert.Equal(t, *updatedUser, editedUser)

	// Print the edited user
	fmt.Println("Edited User:", editedUser)
}

func TestDeleteUser(t *testing.T) {
	userRepo := new(mocks.UserInterfaceRepo)
	userUseCase := useCaseUser{
		userRepo: userRepo,
	}

	userEmail := "john.doe@example.com"

	// Mock the DeleteUser method
	userRepo.On("DeleteUser", userEmail).Return(nil, nil)

	result, err := userUseCase.DeleteUser(userEmail)

	assert.NoError(t, err)
	assert.Nil(t, result)

	// Print a message indicating the user deletion
	fmt.Println("User deleted:", userEmail)
}
