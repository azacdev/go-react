package models

type Entry struct {
	ID         int64    `json:"id" gorm:"primaryKey"`
	Dish       string   `json:"dish" gorm:"not null"`
	Fat        *float64 `json:"fat"`
	Ingredient *string  `json:"ingredient"`
	Calories   *string  `json:"calories"`
}
