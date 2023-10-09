package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/slilp/go-auth/adapter/postgres"
	handler "github.com/slilp/go-auth/handler/account"
)

func main() {
	router := initGin()
	initApplication(router)
	server := runServer(router)
	shutdownServer(server)
}

func initApplication(router *gin.Engine) {
	// logger := monitoring.Logger()

	// healthPath := fmt.Sprintf("%s/health", config.HTTP.Prefix)
	router.GET("/health", func(ctx *gin.Context) {
		// ctx.AbortWithStatusJSON(200, gin.H{
		// 	"statusCode": 200, "message": "OK",
		// })
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db, err := postgres.Initialize()
	if err != nil {
		panic(fmt.Errorf("failed to create database connection: %w", err))
	}

	handler.AccountInitialize(db, &router.RouterGroup)

}

func initGin() *gin.Engine {

	if "production" != "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(
		gin.Recovery(),
	)

	return router
}

func runServer(router *gin.Engine) *http.Server {

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3000),
		Handler: router,
	}
	fmt.Printf("Server is running on port: %s\n", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatal("ListenAndServe")
		}
	}()

	return server
}

func shutdownServer(server *http.Server) {

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()

	stop()
	// log.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown")
	}
}
