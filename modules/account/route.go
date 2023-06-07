package account

import (
	"crud/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouterAccount struct {
	AccountRequestHandler RequestHandlerAccount
}

func NewRouter(
	dbCrud *gorm.DB,
) RouterAccount {
	return RouterAccount{AccountRequestHandler: NewAccountRequestHandler(
		dbCrud,
	)}
}

func (r RouterAccount) Handle(router *gin.Engine) {
	basepath := "/account"
	account := router.Group(basepath)

	account.POST("/",
		r.AccountRequestHandler.CreateAccount,
	)

	account.GET("/:id",
		middleware.AuthMiddleware(),
		r.AccountRequestHandler.GetAccountById,
	)

	account.PATCH("/:id",
		r.AccountRequestHandler.UpdateAccount,
	)

	account.DELETE("/:username",
		r.AccountRequestHandler.DeleteAccount,
	)

	account.POST("/login",
		r.AccountRequestHandler.LoginAccount,
	)

	account.PUT("/verify/:id",
		r.AccountRequestHandler.VerifyAccount,
	)

	account.PUT("/activate/:id",
		r.AccountRequestHandler.ActivateAccount,
	)
}
