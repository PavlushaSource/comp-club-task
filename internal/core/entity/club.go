package entity

import (
	"time"
)

type ClubInfo struct {
	NumberTables int
	OpenTime     time.Time
	CloseTime    time.Time
	PriceHour    int
}
