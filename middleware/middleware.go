package middleware

import (
	"github.com/gin-gonic/gin"
)

type AuthMiddleware interface {
	Authorized() gin.HandlerFunc
	Authentication() gin.HandlerFunc
	Refresh() gin.HandlerFunc
}
