package dtos

type CreateItemReq struct {
	ItemName     string `json:"itemName"`
	Description  string `json:"description"`
	Points       int    `json:"points"`
	NeededAmount int    `json:"neededAmount"`
	EncodedImage string `json:"encodedImage"`
}

type GetItemReq struct {
	ItemID string `param:"id"`
}

type UploadItemImageParams struct {
	ItemID    string
	FileType  string
	ImageData []byte
}
