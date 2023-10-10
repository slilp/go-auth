package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slilp/go-auth/middleware"
	service "github.com/slilp/go-auth/service/account"
	utils "github.com/slilp/go-auth/utils"
	"github.com/spf13/viper"
)

func accountServer(router *gin.RouterGroup, s service.AccountService, middlewareAuth middleware.AuthMiddleware) {
	handlerRoute := NewAccountHttpHandler(s)
	authGroup := router.Group("/auth")
	authGroup.POST("/register", handlerRoute.Register)
	authGroup.POST("/sign-in", handlerRoute.SignIn)
	authGroup.GET("/refresh-token", middlewareAuth.Refresh(), handlerRoute.RefreshToken)

	accountGroup := router.Group("/account", middlewareAuth.Authentication())
	accountGroup.GET("", handlerRoute.GetAccountInfo)
	accountGroup.DELETE("", handlerRoute.DeleteAccount)
	accountGroup.PATCH("", handlerRoute.UpdateAccount)
}

func NewAccountHttpHandler(service service.AccountService) handler {
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

func (h handler) SignIn(ctx *gin.Context) {

	var form service.SignInDto
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	accInfo, err := h.service.SignIn(form)
	if err != nil {
		utils.ReturnError(ctx, err)
		return
	}

	accessTokenChan := make(chan string)
	refreshTokenChan := make(chan string)

	go func() {
		tempAccessToken := utils.GenerateJwtToken(accInfo, viper.GetString("jwt.accessToken"), 10*time.Minute)
		accessTokenChan <- tempAccessToken
	}()

	go func() {
		tempRefreshToken := utils.GenerateJwtToken(accInfo, viper.GetString("jwt.refreshToken"), 24*time.Hour)
		refreshTokenChan <- tempRefreshToken
	}()

	accessToken := <-accessTokenChan
	refreshToken := <-refreshTokenChan

	ctx.JSON(http.StatusOK, gin.H{"account": accInfo, "accessToken": accessToken, "refreshToken": refreshToken})
}

func (h handler) RefreshToken(ctx *gin.Context) {
	session := ctx.MustGet("account").(service.AccountInfo)

	accessToken := utils.GenerateJwtToken(session, viper.GetString("jwt.accessToken"), 10*time.Minute)
	refreshToken := utils.GenerateJwtToken(session, viper.GetString("jwt.refreshToken"), 24*time.Hour)

	ctx.JSON(http.StatusOK, gin.H{"account": session, "accessToken": accessToken, "refreshToken": refreshToken})

}

func (h handler) GetAccountInfo(ctx *gin.Context) {
	session := ctx.MustGet("account").(service.AccountInfo)
	srvRes, err := h.service.GetAccount(session.Username)
	if err != nil {
		utils.ReturnError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, srvRes)
}

func (h handler) DeleteAccount(ctx *gin.Context) {
	session := ctx.MustGet("account").(service.AccountInfo)
	err := h.service.DeleteAccount(session.ID)
	if err != nil {
		utils.ReturnError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (h handler) UpdateAccount(ctx *gin.Context) {

	session := ctx.MustGet("account").(service.AccountInfo)
	req := service.UpdateAccountDto{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.BadRequest(err.Error()))
		return
	}
	err := h.service.UpdateAccount(session.ID, req)
	if err != nil {
		utils.ReturnError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
