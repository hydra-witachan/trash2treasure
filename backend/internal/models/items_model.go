package models

type Item struct {
	BaseModel

	AuthorID        string `gorm:"column:author_id" json:"authorId"`
	CategoryID      string `gorm:"column:category_id" json:"categoryId"`
	SubCategoryID   string `gorm:"column:sub_category_id" json:"subCategoryId"`
	AuthorName      string `gorm:"column:author_name" json:"authorName"`
	ItemName        string `gorm:"column:item_name" json:"itemName"`
	Description     string `gorm:"column:description" json:"description"`
	Points          int    `gorm:"column:points" json:"points"`
	ImageURL        string `gorm:"column:image_url" json:"imageUrl"`
	NeededAmount    int    `gorm:"column:needed_amount" json:"neededAmount"`
	FullfiledAmount int    `gorm:"column:fullfiled_amount" json:"fullfiledAmount"`
}
