package entity

import "time"

type Account struct {
	ID       uint      `gorm:"primary_key"`
	Username string    `gorm:"column:username"`
	Password string    `gorm:"column:password"`
	RoleID   int       `gorm:"column:role_id"`
	Verified string    `gorm:"column:verified"`
	Active   string    `gorm:"column:active"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

func (Account) TableName() string {
	return "account"
}

type Role struct {
	ID       int    `gorm:"primary_key"`
	RoleName string `gorm:"column:role_name"`
}
