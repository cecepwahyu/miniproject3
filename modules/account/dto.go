package account

import (
	"crud/dto"
	"crud/entity"
	"time"
)

type AccountParam struct {
	ID       uint      `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	RoleID   int       `json:"role_id"`
	Verified string    `json:"verified"`
	Active   string    `json:"active"`
	CreateAt time.Time `json:"created_at"`
	UpdateAt time.Time `json:"updated_at"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data AccountParam `json:"data"`
}

type FindAccount struct {
	dto.ResponseMeta
	Data entity.Account `json:"data"`
}

type LoginAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
