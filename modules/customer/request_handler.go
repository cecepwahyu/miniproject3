package customer

import (
	"crud/dto"
	"crud/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandlerCustomer struct {
	ctr ControllerCustomer
}

func NewCustomerRequestHandler(
	dbCrud *gorm.DB,
) RequestHandlerCustomer {
	return RequestHandlerCustomer{
		ctr: controllerCustomer{
			customerUseCase: useCaseCustomer{
				customerRepo: repository.NewCustomer(dbCrud),
			},
		}}
}

func (h RequestHandlerCustomer) CreateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.CreateCustomer(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) GetCustomerById(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	customerid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultErrorResponse())
		return
	}

	//batas tambahan fitur yang bisa diblock
	res, err := h.ctr.GetCustomerById(uint(customerid))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) UpdateCustomer(c *gin.Context) {
	request := CustomerParam{}
	err := c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	customerId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := h.ctr.UpdateCustomer(request, uint(customerId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h RequestHandlerCustomer) DeleteCustomer(c *gin.Context) {
	email := c.Param("email")
	res, err := h.ctr.DeleteCustomer(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
