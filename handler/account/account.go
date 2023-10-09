package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/slilp/go-auth/service/account"
	utils "github.com/slilp/go-auth/util"
)

func accountServer(router *gin.RouterGroup, s service.AccountService) {
	handlerRoute := accountHttpHandler(s)
	authGroup := router.Group("/auth")
	authGroup.POST("/register", handlerRoute.Register)
	authGroup.POST("/sign-in", handlerRoute.SignIn)
	authGroup.GET("/refresh-token", handlerRoute.RefreshToken)

	accountGroup := router.Group("/account")
	accountGroup.GET("", handlerRoute.GetAccountInfo)
}

func accountHttpHandler(service service.AccountService) handler {
	return handler{service: service}
}

type handler struct {
	service service.AccountService
}

func (h handler) Register(ctx *gin.Context) {
	req := service.CreateAccountDto{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BadRequest(err.Error()))
		return
	}

	accountRes, err := h.service.CreateAccount(req)
	if err != nil {
		utils.ReturnError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, accountRes)
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

	srvRes, err := h.service.GetAccount("slil.pua@gmail.com")
	if err != nil {
		utils.ReturnError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, srvRes)
}
