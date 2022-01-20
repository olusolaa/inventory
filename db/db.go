package db

import (
	"github.com/olusolaa/inventry/models"
)

// inventory service
type DB interface {
	// inventory
	GetCategory(id int64) (*models.Category, error)
	GetCategories() ([]*models.Category, error)
	CreateCategory(inventory *models.Category) (*models.Category, error)
	UpdateCategory(inventory *models.Category) (*models.Category, error)
	DeleteCategory(id int64) (*models.Category, error)

	// inventory item
	GetItem(id int64) (*models.Item, error)
	GetItems() ([]*models.Item, error)
	CreateItem(inventoryItem *models.Item) (*models.Item, error)
	UpdateItem(inventoryItem *models.Item) (*models.Item, error)
	DeleteItem(id int64) (*models.Item, error)
	GetCategoryItems(categoryId int64) ([]*models.Item, error)
}
