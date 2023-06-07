package customer

import (
	"crud/entity"
	"crud/repository"
	"time"
)

type UseCaseCustomer interface {
	CreateCustomer(customer CustomerParam) (entity.Customer, error)
	GetCustomerById(id uint) (entity.Customer, error)
	UpdateCustomer(customer CustomerParam, id uint) (any, error)
	DeleteCustomer(email string) (any, error)
}

type useCaseCustomer struct {
	customerRepo repository.CustomerInterfaceRepo
}

func (uc useCaseCustomer) CreateCustomer(customer CustomerParam) (entity.Customer, error) {
	var newCustomer *entity.Customer

	newCustomer = &entity.Customer{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	_, err := uc.customerRepo.CreateCustomer(newCustomer)
	if err != nil {
		return *newCustomer, err
	}
	return *newCustomer, nil
}

func (uc useCaseCustomer) GetCustomerById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	customer, err := uc.customerRepo.GetCustomersById(id)
	return customer, err
}

func (uc useCaseCustomer) UpdateCustomer(customer CustomerParam, id uint) (any, error) {
	var editCustomer *entity.Customer
	editCustomer = &entity.Customer{
		ID:        id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	_, err := uc.customerRepo.UpdateCustomer(editCustomer)
	if err != nil {
		return *editCustomer, err
	}
	return *editCustomer, nil
}

func (uc useCaseCustomer) DeleteCustomer(email string) (any, error) {
	_, err := uc.customerRepo.DeleteCustomer(email)
	return nil, err
}
