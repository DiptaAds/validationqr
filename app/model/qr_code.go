package model

type QRCode struct {
	Product
	UUID         string `json:"uuid" bson:"uuid"`
	ActiveStatus bool   `json:"activestatus" bson:"activestatus"`
	Amount       int    `json:"amount" bson:"amount"`
	BaseModel
}
