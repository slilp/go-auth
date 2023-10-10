package middleware

import (
	"net/http"
	"strings"

	service "github.com/slilp/go-auth/service/account"
	"github.com/slilp/go-auth/utils"

	"github.com/gin-gonic/gin"
)

type jwtMiddleware struct {
	accessTokenSecret  string
	refreshTokenSecret string
}

func NewJwtMiddleware(accessTokenSecret string, refreshTokenSecret string) AuthMiddleware {
	return jwtMiddleware{accessTokenSecret, refreshTokenSecret}
}

func (j jwtMiddleware) Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")

		if headerToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")
		var account service.AccountInfo
		if err := utils.ValidateToken(token, &account, j.accessTokenSecret); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("account", account)
	}
}

func (j jwtMiddleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")

		if headerToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")
		var account service.AccountInfo
		if err := utils.ValidateToken(token, &account, j.accessTokenSecret); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("account", account)
	}
}

func (j jwtMiddleware) Refresh() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.Request.Header.Get("Authorization")

		if headerToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")
		var account service.AccountInfo
		if err := utils.ValidateToken(token, &account, j.refreshTokenSecret); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Set("account", account)

	}
}
