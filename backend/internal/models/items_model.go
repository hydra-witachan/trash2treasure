package models

type Item struct {
	BaseModel

	AuthorID        string `gorm:"column:author_id" json:"authorId"`
	AuthorName      string `gorm:"column:author_name" json:"authorName"`
	ItemName        string `gorm:"column:item_name" json:"itemName"`
	Description     string `gorm:"column:description" json:"description"`
	Points          int    `gorm:"column:points" json:"points"`
	ImageURL        string `gorm:"column:image_url" json:"imageUrl"`
	NeededAmount    int    `gorm:"column:needed_amount" json:"neededAmount"`
	FullfiledAmount int    `gorm:"column:fullfiled_amount" json:"fullfiledAmount"`
	Category		string `gorm:"column:category" json:"category"`
	SubCategory		string `gorm:"column:sub_category" json:"subCategory"`
}
