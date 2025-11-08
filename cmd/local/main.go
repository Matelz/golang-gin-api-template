package main

import (
	"github.com/gin-gonic/gin"

	"devmaua.com/devbank/internal/account"
	"devmaua.com/devbank/internal/user"
)

func main() {
	memoryUserRepo := user.NewMemoryRepository()

	userService := user.NewService(memoryUserRepo)

	memoryAccountRepo := account.NewMemoryRepository()

	accountService := account.NewService(memoryAccountRepo, *userService)

	// Initialize handlers
	userHandler := user.NewHandler(userService)
	accountHandler := account.NewHandler(accountService)

	r := gin.Default()

	// Account routes
	r.POST("/api/v1/account/deposit", accountHandler.DepositAmount)
	r.POST("/api/v1/account/withdraw", accountHandler.WithdrawAmount)
	r.POST("/api/v1/account", accountHandler.CreateAccount)
	r.GET("/api/v1/account/:id", accountHandler.GetAccount)

	// User routes
	r.POST("/api/v1/user", userHandler.CreateUser)
	r.GET("/api/v1/user/:id", userHandler.GetUserByID)

	r.Run() // listens on 0.0.0.0:8080 by default
}
