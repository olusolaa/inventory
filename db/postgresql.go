package db

import (
	"fmt"
	"github.com/olusolaa/inventry/models"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

// PostgresDB implements the DB interface
type PostgresDB struct {
	DB *gorm.DB
}

func (psql *PostgresDB) GetCategory(id int64) (*models.Category, error) {
	var category models.Category
	if err := psql.DB.Preload("Items").Where("id = ?", id).First(&category).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get category")
	}
	return &category, nil
}

func (psql *PostgresDB) GetCategories() ([]*models.Category, error) {
	var categories []*models.Category
	if err := psql.DB.Find(&categories).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get categories")
	}
	return categories, nil
}

func (psql *PostgresDB) CreateCategory(inventory *models.Category) (*models.Category, error) {
	if err := psql.DB.Create(inventory).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create category")
	}
	return inventory, nil
}

func (psql *PostgresDB) UpdateCategory(inventory *models.Category) (*models.Category, error) {
	if err := psql.DB.Save(inventory).Error; err != nil {
		return nil, errors.Wrap(err, "failed to update category")
	}
	return inventory, nil
}

func (psql *PostgresDB) DeleteCategory(id int64) (*models.Category, error) {
	var category models.Category
	if err := psql.DB.Where("id = ?", id).Delete(&category).Error; err != nil {
		return nil, errors.Wrap(err, "failed to delete category")
	}
	if err := psql.DB.Where("category_id = ?", id).Delete(&models.Item{}).Error; err != nil {
		return nil, errors.Wrap(err, "failed to delete category items")
	}
	return &category, nil
}

func (psql *PostgresDB) GetItem(id int64) (*models.Item, error) {
	var item models.Item
	if err := psql.DB.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get item")
	}
	return &item, nil
}

func (psql *PostgresDB) GetItems() ([]*models.Item, error) {
	var items []*models.Item
	if err := psql.DB.Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get items")
	}
	return items, nil
}

func (psql *PostgresDB) CreateItem(inventoryItem *models.Item) (*models.Item, error) {
	if err := psql.DB.Create(inventoryItem).Error; err != nil {
		return nil, errors.Wrap(err, "failed to create item")
	}
	return inventoryItem, nil
}

func (psql *PostgresDB) UpdateItem(inventoryItem *models.Item) (*models.Item, error) {
	var item models.Item
	psql.DB.First(&item)
	item.Name = inventoryItem.Name
	item.CategoryId = inventoryItem.CategoryId
	item.Price = inventoryItem.Price
	item.Quantity = inventoryItem.Quantity
	item.Description = inventoryItem.Description

	if err := psql.DB.Save(&item).Error; err != nil {
		return nil, errors.Wrap(err, "failed to update item")
	}
	return &item, nil
}

func (psql *PostgresDB) DeleteItem(id int64) (*models.Item, error) {
	var item models.Item
	if err := psql.DB.Where("id = ?", id).Delete(&item).Error; err != nil {
		return nil, errors.Wrap(err, "failed to delete item")
	}
	return &item, nil
}

func (psql *PostgresDB) GetCategoryItems(id int64) ([]*models.Item, error) {
	//retrieve all items for a category
	var items []*models.Item
	if err := psql.DB.Where("category_id = ?", id).Find(&items).Error; err != nil {
		return nil, errors.Wrap(err, "failed to get items")
	}
	fmt.Printf("items: %+v\n", items)
	return items, nil
}

// Init initializes the database connection
func (psql *PostgresDB) Init() {
	psqlInfo := os.Getenv("DATABASE_URL")
	if psqlInfo == "" {
		psqlInfo = "postgres://postgres:postgres@localhost:5431/inventory?sslmode=disable"
	}
	DBSession, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(errors.Wrap(err, "Unable to connect to Postgresql database"))
	}
	psql.DB = DBSession
	err = psql.DB.AutoMigrate(&models.Item{}, models.Category{})
	if err != nil {
		panic(errors.Wrap(err, "Unable to migrate Postgresql database"))
	}
}
