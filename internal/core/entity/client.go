package entity

import (
	"time"
)

type Client struct {
	Status   int
	Table    int
	SeatTime time.Time
}
