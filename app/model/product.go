package model

type Product struct {
	BaseModel
	Name string `json:"name" bson:"name"`
}
