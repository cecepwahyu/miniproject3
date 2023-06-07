package account

import (
	"crud/dto"
	"crud/entity"
)

type AccountParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	RoleID   int    `json:"role_id"`
	Verified string `json:"verified"`
	Active   string `json:"active"`
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
