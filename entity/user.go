package entity

import "time"

type User struct {
	ID       uint      `gorm:"primary_key"`
	Name     string    `gorm:"column:name"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`
	CreateAt time.Time `gorm:"column:created_at"`
	UpdateAt time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "user"
}
