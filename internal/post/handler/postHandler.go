package handler

import (
	"Project1/internal/post/model"
	"Project1/internal/post/service"
	handler "Project1/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "image/jpeg"
	_ "image/png"
	"net/http"
	"path/filepath"
	"strconv"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) CreatePost(c *gin.Context) {
	var post model.Post
	err := c.ShouldBind(&post)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		handler.NewErrorResponse(c, 500, "invalid file")
		return
	}

	extension := filepath.Ext(file.Filename)
	newFileName := uuid.New().String() + extension

	if err := c.SaveUploadedFile(file, "uploads/images/"+newFileName); err != nil {
		handler.NewErrorResponse(c, 500, "failed storing image")
		return
	}
	imgPath := "uploads/images/" + newFileName
	post.ImagePath = &imgPath

	newPost, err := h.service.CreatePost(&post)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"data": newPost,
	})
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, "something went wrong")
		return
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func (h *PostHandler) GetPostByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, "invalid id")
		return
	}

	post, err := h.service.GetPostByID(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, "something went wrong")
		return
	}

	c.JSON(200, post)
}

func (h *PostHandler) DeletePost(c *gin.Context) {
	postId := c.Param("id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeletePostById(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, "does not implement service")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}

func (h *PostHandler) UpdatePost(c *gin.Context) {
	postId := c.Param("id")
	id, err := strconv.Atoi(postId)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var updatedPost model.Post
	err = c.BindJSON(&updatedPost)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.service.UpdatePost(id, &updatedPost)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, "Failed to updated post")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post updated successfully"})
}
