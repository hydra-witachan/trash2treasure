package models

type Category struct {
	BaseModel

	Name	string	`gorm:"column:name" json:"name"`
}