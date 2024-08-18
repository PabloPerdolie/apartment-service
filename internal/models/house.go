package models

import (
	"database/sql"
	"time"
)

type House struct {
	Id        int32
	Address   string
	Year      int32
	Developer string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
