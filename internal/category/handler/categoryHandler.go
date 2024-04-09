package handler

import (
	"Project1/internal/category/model"
	"Project1/internal/category/service"
	handler "Project1/pkg/response"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) PostCategory(c *gin.Context) {
	var category model.Category
	err := c.BindJSON(&category)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	newCategory, err := h.service.CreateCategory(&category)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
		return
	}
	c.JSON(200, gin.H{
		"data": newCategory,
	})

}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	Categories, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(500, "Something went wrong")
		return
	}

	c.JSON(200, gin.H{
		"data": Categories,
	})
}

func (h *CategoryHandler) GetOneCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, "invalid id")
		return
	}

	category, err := h.service.GetOneCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "something went wrong")
		return
	}
	c.JSON(200, gin.H{
		"data": category,
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.DeleteCategoryById(id)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusUnauthorized, "does not get service")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

func (h *CategoryHandler) UpdateCategoryPls(c *gin.Context) {
	categoryId := c.Param("id")
	id, err := strconv.Atoi(categoryId)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var category model.Category
	if err := c.BindJSON(&category); err != nil {
		handler.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.UpdateCategory(id, &category)
	if err != nil {
		handler.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}
