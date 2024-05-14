package entity

import (
	"time"
)

const (
	ClientCome = iota + 1
	ClientSeatTable
	ClientWaiting
	ClientLeft
	ClientLeftClub = iota + 7
	ClientSeatAfterWaiting
	EventFailed
)

type Event struct {
	Time        time.Time
	ID          int
	ClientName  string
	TableNumber int
	Error       error
}
