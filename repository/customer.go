package repository

import (
	"crud/entity"
	"gorm.io/gorm"
)

type Customer struct {
	db *gorm.DB
}

func NewCustomer(dbCrud *gorm.DB) Customer {
	return Customer{
		db: dbCrud,
	}

}

type CustomerInterfaceRepo interface {
	CreateCustomer(customer *entity.Customer) (*entity.Customer, error)
	GetCustomersById(id uint) (entity.Customer, error)
	UpdateCustomer(customer *entity.Customer) (any, error)
	DeleteCustomer(email string) (any, error)
}

func (repo Customer) CreateCustomer(customer *entity.Customer) (*entity.Customer, error) {
	err := repo.db.Model(&entity.Customer{}).Create(customer).Error
	return customer, err
}

func (repo Customer) GetCustomersById(id uint) (entity.Customer, error) {
	var customer entity.Customer
	repo.db.First(&customer, "id = ? ", id)
	return customer, nil
}

// UpdateUser multiple fields
func (repo Customer) UpdateCustomer(customer *entity.Customer) (any, error) {
	err := repo.db.Model(&entity.Customer{}).
		Save(customer).Error
	return nil, err
}

// DeleteUser by Id and email
func (repo Customer) DeleteCustomer(email string) (any, error) {
	err := repo.db.Model(&entity.Customer{}).
		Where("email = ?", email).
		Delete(&entity.Customer{}).
		Error
	return nil, err
}
