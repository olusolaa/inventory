package server

import (
	"github.com/gin-gonic/gin"
	"github.com/olusolaa/inventry/models"
	"log"
	"net/http"
	"strconv"
)

// GetItems godoc
// @Summary Get all items
// @Description Get all items
// @Tags items
// @Produce json
// @Success 200 {array} []models.Item
// @Failure 500 {object} models.Error
// @Router /api/v1/items [get]
func (s *Server) GetItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		items, err := s.DB.GetItems()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}
}

// GetItem godoc
// @Summary Get item by id
// @Description Get item by id
// @Tags items
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 404 {object} models.Error
// @Router /api/v1/item/{id} [get]
func (s *Server) GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		item, err := s.DB.GetItem(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, item)
	}
}

// CreateItem godoc
// @Summary Create item
// @Description Create item
// @Tags items
// @Accept json
// @Produce json
// @Param item body models.ItemDto true "Item"
// @Success 200 {object} models.Item
// @Failure 400 {object} models.Error
// @Router /api/v1/items [post]
func (s *Server) CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		if err := s.decode(c, &item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data, err := s.DB.CreateItem(&item)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Item added successfully!", "data": data})
	}
}

// UpdateItem godoc
// @Summary Update item
// @Description Update item
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Param item body models.ItemDto true "Item"
// @Success 200 {object} models.Item
// @Failure 400 {object} models.Error
// @Router /api/v1/items/{id} [put]
func (s *Server) UpdateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var item models.Item
		if err := s.decode(c, &item); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		data, err := s.DB.UpdateItem(&item)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully!", "data": data})
	}
}

// DeleteItem godoc
// @Summary Delete item
// @Description Delete item
// @Tags items
// @Accept json
// @Produce json
// @Param id path int true "Item ID"
// @Success 200 {object} models.Item
// @Failure 400 {object} models.Error
// @Router /api/v1/items/{id} [delete]
func (s *Server) DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		data, err := s.DB.DeleteItem(id)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully!", "data": data})
	}
}
