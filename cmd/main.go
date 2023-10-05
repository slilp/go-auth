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
)

var server *http.Server
var router *gin.Engine

func main() {
	router := initGin()
	initApplication(router)
	runServer()
	shutdownServer()
}

func initApplication(router *gin.Engine) {
	// logger := monitoring.Logger()

	// healthPath := fmt.Sprintf("%s/health", config.HTTP.Prefix)
	router.GET("health", func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(200, gin.H{
			"statusCode": 200, "message": "OK",
		})
	})
}

func initGin() *gin.Engine {

	if "production" == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router = gin.New()
	router.Use(
		gin.Recovery(),
	)

	return router
}

func runServer() *http.Server {

	server = &http.Server{
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

func shutdownServer() {

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
