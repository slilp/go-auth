package handler

import (
	"github.com/slilp/go-auth/middleware"
	repository "github.com/slilp/go-auth/repository/account"
	service "github.com/slilp/go-auth/service/account"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountInitialize(
	db *gorm.DB,
	group *gin.RouterGroup,
	middlewareAuth middleware.AuthMiddleware,
) {
	repo := repository.NewAccountRepository(db)
	service := service.NewAccountService(
		repo,
	)
	accountServer(group, service, middlewareAuth)
}
