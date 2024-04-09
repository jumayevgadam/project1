package routes

import (
	"Project1/internal/category/handler"
	"Project1/internal/category/repository"
	"Project1/internal/category/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitCategoryRoutes(router *gin.RouterGroup, DB *sqlx.DB) {
	categoryRepo := repository.NewCategoryRepository(DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	categoryRoutes := router.Group("/category")
	{
		categoryRoutes.GET("/GetAll", categoryHandler.GetAllCategories)
		categoryRoutes.GET("/:id", categoryHandler.GetOneCategory)
		categoryRoutes.POST("/post", categoryHandler.PostCategory)
		categoryRoutes.PUT("/:id", categoryHandler.UpdateCategoryPls)
		categoryRoutes.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
