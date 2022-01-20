package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/inventry/models"
	"log"
	"net/http"
	"strconv"
)

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category
// @Tags category
// @Accept  json
// @Produce  json
// @Param category body models.CategoryDto true "Category"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.Error
// @Router /api/v1/categories [post]
func (s *Server) CreateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		if err := s.decode(c, &category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data, err := s.DB.CreateCategory(&category)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Category added successfully!", "data": data})
	}
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category
// @Tags category
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body models.CategoryDto true "Category"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.Error
// @Router /api/v1/categories/{id} [put]
func (s *Server) UpdateCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		var category models.Category
		if err := s.decode(c, &category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data, err := s.DB.UpdateCategory(&category)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully!", "data": data})
	}
}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category
// @Tags category
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Failure 500 {object} models.Error
// @Router /api/v1/categories/{id} [delete]
func (s *Server) DeleteCategory() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		data, err := s.DB.DeleteCategory(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully!", "data": data})
	}
}

// GetCategoryItems godoc
// @Summary Get category items
// @Description Get category items
// @Tags category
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} []models.Item
// @Failure 500 {object} models.Error
// @Router /api/v1/categories/{id}/items [get]
func (s *Server) GetCategoryItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data, err := s.DB.GetCategoryItems(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(id)
		c.JSON(http.StatusOK, gin.H{"message": "Category items retrieved successfully!", "data": data})
	}
}

// GetCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags category
// @Produce  json
// @Success 200 {object} []models.Category
// @Failure 500 {object} models.Error
// @Router /api/v1/categories [get]
func (s *Server) GetCategories() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := s.DB.GetCategories()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Categories retrieved successfully!", "data": data})
	}
}
