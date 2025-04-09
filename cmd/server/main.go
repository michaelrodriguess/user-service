package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/michaelrodriguess/user-service/internal/handler"
	"github.com/michaelrodriguess/user-service/internal/model"
	"github.com/michaelrodriguess/user-service/internal/repository"
	"github.com/michaelrodriguess/user-service/internal/service"
	"github.com/michaelrodriguess/user-service/pkg/client"
	mysqldb "github.com/michaelrodriguess/user-service/pkg/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  .env not found, using default environment variables")
	}

	mysqldb.Init()

	mysqldb.DB.AutoMigrate(&model.User{})

	authClient := client.NewAuthClient()
	repo := repository.NewUserRepository(mysqldb.DB)
	service := service.NewUserService(repo, authClient)
	handler := handler.NewUserHandler(service)

	r := gin.Default()

	r.POST("/users", handler.CreateUserHandler)
	r.GET("/admin-users", handler.GetAllAdminsUser)
	r.GET("/users", handler.GetAllUsers)
	r.DELETE("/users", handler.DeleteUserHandler)
	r.PATCH("/users", handler.UpdateUserHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("✅ User service running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("❌ Server failed to start:", err)
	}
}
