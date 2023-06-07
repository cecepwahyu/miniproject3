package customer

import (
	"crud/dto"
	"crud/entity"
)

type CustomerParam struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

//type CustomerParam struct {
//	FirstName string `json:"first_name"`
//	LastName  string `json:"last_name"`
//	Email     string `json:"email"`
//	Avatar    string `json:"avatar"`
//}

type SuccessCreate struct {
	dto.ResponseMeta
	Data CustomerParam `json:"data"`
}

type FindCustomer struct {
	dto.ResponseMeta
	Data entity.Customer `json:"data"`
}
