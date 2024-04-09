package routes

import (
	"Project1/internal/users/handler"
	"Project1/internal/users/repository"
	"Project1/internal/users/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitUserRoutes(router *gin.RouterGroup, DB *sqlx.DB) {
	userRepo := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	userRoutes := router.Group("/users")
	userRoutes.POST("/sign-in", userHandler.SignIn)
	userRoutes.POST("/sign-up", userHandler.SignUp)
	userRoutes.DELETE("/:id", userHandler.DeleteUser)
}
