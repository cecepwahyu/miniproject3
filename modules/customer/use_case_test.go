package customer

import (
	"crud/entity"
	"crud/repository/mocks"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestCreateCustomer(t *testing.T) {
	customerRepo := new(mocks.CustomerInterfaceRepo)
	customerUseCase := useCaseCustomer{
		customerRepo: customerRepo,
	}

	customerParam := CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
	}

	newCustomer := &entity.Customer{
		FirstName: customerParam.FirstName,
		LastName:  customerParam.LastName,
		Email:     customerParam.Email,
		Avatar:    customerParam.Avatar,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	// Mock the CreateCustomer method
	customerRepo.On("CreateCustomer", newCustomer).Return(newCustomer, nil)

	createdCustomer, err := customerUseCase.CreateCustomer(customerParam)

	assert.NoError(t, err)
	assert.Equal(t, *newCustomer, createdCustomer)

	// Print the created customer
	fmt.Println("Created Customer:", createdCustomer)
}

func TestDeleteCustomer(t *testing.T) {
	customerRepo := new(mocks.CustomerInterfaceRepo)
	customerUseCase := useCaseCustomer{
		customerRepo: customerRepo,
	}

	email := "john.doe@example.com"

	// Mock the DeleteCustomer method
	customerRepo.On("DeleteCustomer", email).Return(nil, nil)

	result, err := customerUseCase.DeleteCustomer(email)

	assert.NoError(t, err)
	assert.Nil(t, result)

	// Print a message indicating the customer deletion
	fmt.Println("Customer deleted:", email)
}

func TestGetCustomerById(t *testing.T) {
	customerRepo := new(mocks.CustomerInterfaceRepo)
	customerUseCase := useCaseCustomer{
		customerRepo: customerRepo,
	}

	customerID := uint(1)

	expectedCustomer := entity.Customer{
		ID:        customerID,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	// Mock the GetCustomersById method
	customerRepo.On("GetCustomersById", customerID).Return(expectedCustomer, nil)

	retrievedCustomer, err := customerUseCase.GetCustomerById(customerID)

	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, retrievedCustomer)

	// Print the retrieved customer
	fmt.Println("Retrieved Customer:", retrievedCustomer)
}

func TestUpdateCustomer(t *testing.T) {
	customerRepo := new(mocks.CustomerInterfaceRepo)
	customerUseCase := useCaseCustomer{
		customerRepo: customerRepo,
	}

	customerParam := CustomerParam{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Avatar:    "avatar.png",
	}

	customerID := uint(1)

	editedCustomer := &entity.Customer{
		ID:        customerID,
		FirstName: customerParam.FirstName,
		LastName:  customerParam.LastName,
		Email:     customerParam.Email,
		Avatar:    customerParam.Avatar,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	// Mock the UpdateCustomer method
	customerRepo.On("UpdateCustomer", mock.AnythingOfType("*entity.Customer")).Return(editedCustomer, nil)

	updatedCustomer, err := customerUseCase.UpdateCustomer(customerParam, customerID)

	assert.NoError(t, err)
	assert.Equal(t, *editedCustomer, updatedCustomer)

	// Print the updated customer
	fmt.Println("Updated Customer:", updatedCustomer)
}
