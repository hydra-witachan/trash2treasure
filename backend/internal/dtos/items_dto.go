package dtos

type CreateItemReq struct {
	ItemName     string `json:"itemName"`
	Description  string `json:"description"`
	Points       int    `json:"points"`
	NeededAmount int    `json:"neededAmount"`
	ImageURL     string `json:"imageUrl"`
}

type GetItemReq struct {
	ItemID string `param:"id"`
}
