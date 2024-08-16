package models

type Flat struct {
	Id          int
	HouseId     int
	Price       int
	Rooms       int8
	Status      string
	ModeratorId string
}
