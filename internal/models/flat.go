package models

type Flat struct {
	Id          int32
	HouseId     int32
	Price       int32
	Rooms       int32
	Status      string
	ModeratorId string
}
