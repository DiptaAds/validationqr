package model

type Conveyor struct {
	Motor bool `json:"motor" bson:"motor"`
	Servo bool `json:"servo" bson:"servo"`
	Speed int  `json:"speed" bson:"speed"`
}
