package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thalerngsak/todoapplication/handler"
	"github.com/thalerngsak/todoapplication/middleware"
	"github.com/thalerngsak/todoapplication/model"
	"github.com/thalerngsak/todoapplication/repository"
	"github.com/thalerngsak/todoapplication/service"
	"github.com/thalerngsak/todoapplication/token"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load("local.env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err.Error())
	}

	err = db.AutoMigrate(&model.User{}, &model.Todo{})
	if err != nil {
		return
	}
	tokenMaker := token.NewJWTMaker(os.Getenv("JWT_SECRET_KEY"))
	if err != nil {
		log.Fatalf("Error creating token maker: %s", err.Error())
	}

	userStore := repository.NewUserDB(db)
	userService := service.NewUserService(userStore)
	userHandler := handler.NewUserHandler(userService, tokenMaker)

	todoStore := repository.NewTodoDB(db)
	todoService := service.NewTodoService(todoStore)
	todoHandler := handler.NewTodoHandler(todoService)

	r := gin.New()
	r.Use(gin.Logger())

	r.POST("/api/login", userHandler.Login)

	api := r.Group("/api")
	api.Use(middleware.AuthenticationMiddleware(tokenMaker))

	api.POST("/todos", todoHandler.Create)
	api.PUT("/todos/:id", todoHandler.Update)
	api.PATCH("/todos/:id/done", todoHandler.MarkAsDone)
	api.DELETE("/todos/:id", todoHandler.Delete)
	api.GET("/todos", todoHandler.List)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err = r.Run(":" + port)
	if err != nil {
		log.Fatalf("Error starting server: %s", err.Error())
	}

}
