package routes

import (
	"Project1/internal/post/handler"
	"Project1/internal/post/repository"
	"Project1/internal/post/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitPostRoutes(router *gin.RouterGroup, DB *sqlx.DB) {

	postRepo := repository.NewPostRepository(DB)
	postService := service.NewPostService(postRepo)
	postHandler := handler.NewPostHandler(postService)

	postRoutes := router.Group("/post")
	{
		postRoutes.GET("/", postHandler.GetAllPosts)
		postRoutes.POST("/", postHandler.CreatePost)
		postRoutes.GET("/:id", postHandler.GetPostByID)
		postRoutes.PUT("/:id", postHandler.UpdatePost)
		postRoutes.DELETE("/:id", postHandler.DeletePost)
	}
}
