package models

import "time"

type House struct {
	Id        int
	Address   string
	Year      int16
	Developer string
	CreatedAt time.Time
	UpdatedAt time.Time
}
