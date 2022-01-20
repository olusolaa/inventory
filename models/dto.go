package models

type ItemDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int64  `json:"quantity"`
	CategoryId  int64  `json:"category_id"`
}

type CategoryDto struct {
	Name string `json:"name"`
}

type Error struct {
	Message string `json:"message"`
}
