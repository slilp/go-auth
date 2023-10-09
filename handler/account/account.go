package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/slilp/go-auth/service/account"
)

func accountServer(router *gin.RouterGroup, s service.AccountService) {
	handlerRoute := accountHttpHandler(s)
	authGroup := router.Group("/auth")
	authGroup.POST("/register", handlerRoute.Register)
	authGroup.POST("/sign-in", handlerRoute.SignIn)
	authGroup.GET("/refresh-token", handlerRoute.RefreshToken)

	accountGroup := router.Group("/account")
	accountGroup.GET("", handlerRoute.GetAccountInfo)
	// accountGroup.PATCH("/update-info", PreSignedURLHandler(s))
	// accountGroup.DELETE("", PreSignedURLHandler(s))
}

func accountHttpHandler(service service.AccountService) handler {
	return handler{service: service}
}

type handler struct {
	service service.AccountService
}

func (h handler) Register(ctx *gin.Context) {
	req := service.CreateAccountDto{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	h.service.CreateAccount(req)
}

func (h handler) RefreshToken(ctx *gin.Context) {
	req := service.CreateAccountDto{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	h.service.CreateAccount(req)
}

func (h handler) SignIn(ctx *gin.Context) {
	req := service.CreateAccountDto{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	h.service.CreateAccount(req)
}

func (h handler) GetAccountInfo(ctx *gin.Context) {
	req := service.CreateAccountDto{}
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	h.service.CreateAccount(req)
}
