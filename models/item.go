package models

type Item struct {
	Id          int64  `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	CategoryId  int64  `json:"category_id" gorm:"foreignkey:CategoryId"`
	UserId      int64  `json:"user_id" gorm:"foreignkey:UserId"`
}
