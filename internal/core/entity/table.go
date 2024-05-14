package entity

import (
	"time"
)

type Table struct {
	UsedTime time.Time
	Profit   int
	IsUsing  bool
}
