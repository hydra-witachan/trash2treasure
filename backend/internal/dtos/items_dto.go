package dtos

type CreateItemReq struct {
	Category	  string `json:"category"`
	SubCategory   string `json:"subCategory"`
	ItemName      string `json:"itemName"`
	Description   string `json:"description"`
	PointsPerItem int    `json:"pointsPerItem"`
	NeededAmount  int    `json:"neededAmount"`
	EncodedImage  string `json:"encodedImage"`
}

type GetItemByIDReq struct {
	ItemID string `param:"id"`
}

type GetCollectorItemsReq struct {
	CollectorID string `param:"id"`
}

type GetItemsReq struct {
	SubCategory 	string 	`query:"sub_category"`
	Search			string	`query:"search"`
}

type UploadItemImageParams struct {
	ItemID    string
	FileType  string
	ImageData []byte
}

type DonateItemReq struct {
	ItemID string `param:"id" json:"id"`

	PhoneNumber string `json:"phoneNumber"`
	Address     string `json:"address"`
	Quantity    int    `json:"quantity"`
}
