package main

import (
	"context"
	"log"

	"devmaua.com/devbank/internal/account"
	"devmaua.com/devbank/internal/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	log.Println("Gin cold start")

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

	ginLambda = ginadapter.New(r)

	log.Println("Gin cold start complete")
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
