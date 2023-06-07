package account

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerAccount struct {
	ctr ControllerAccount
}

func NewAccountRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerAccount {
	return RequestHandlerAccount{
		ctr: controllerAccount{
			accountUseCase: useCaseAccount{
				accountRepo: repository.NewAccount(dbCrud),
			},
		}}
}

func (h RequestHandlerAccount) CreateAccount(c *gin.Context) {
	request := AccountParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateAccount(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAccount) GetAccountById(c *gin.Context) {
	request := AccountParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	accountid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	//batas tambahan fitur yang bisa diblock
	res, err := h.ctr.GetAccountById(uint(accountid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAccount) UpdateAccount(c *gin.Context) {
	request := AccountParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	accountId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.UpdateAccount(request, uint(accountId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAccount) DeleteAccount(c *gin.Context) {
	username := c.Param("username")
	res, err := h.ctr.DeleteAccount(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAccount) LoginAccount(c *gin.Context) {
	request := LoginAccount{}
	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	result, err := h.ctr.LoginAccount(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h RequestHandlerAccount) VerifyAccount(c *gin.Context) {
	request := AccountParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	accountId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.VerifyAccount(request, uint(accountId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerAccount) ActivateAccount(c *gin.Context) {
	request := AccountParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	accountId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.ActivateAccount(request, uint(accountId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
