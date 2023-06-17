package dtos

type GetSubCategoriesReq struct {
	CategoryID	string	`query:"category_id"`
}