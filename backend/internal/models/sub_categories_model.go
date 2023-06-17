package models

type SubCategory struct {
	BaseModel

	CategoryID	string	`gorm:"column:category_id" json:"categoryId"`
	Name		string	`gorm:"column:name" json:"name"`
}