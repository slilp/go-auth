package handler

import (
	repository "github.com/slilp/go-auth/repository/account"
	service "github.com/slilp/go-auth/service/account"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AccountInitialize(
	db *gorm.DB,
	group *gin.RouterGroup,
) {
	repo := repository.NewAccountRepository(db)
	service := service.NewAccountService(
		repo,
	)
	accountServer(group, service)
}
