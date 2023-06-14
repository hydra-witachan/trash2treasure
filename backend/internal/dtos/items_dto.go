package dtos

type CreateItemReq struct {
	ItemName      string `json:"itemName"`
	Description   string `json:"description"`
	PointsPerItem int    `json:"pointsPerItem"`
	NeededAmount  int    `json:"neededAmount"`
	EncodedImage  string `json:"encodedImage"`
}

type GetItemByIDReq struct {
	ItemID string `param:"id"`
}

type UploadItemImageParams struct {
	ItemID    string
	FileType  string
	ImageData []byte
}
