package utils

import (
	"time"
)

func EventInWorkingTime(open, close, eventTime time.Time) bool {
	if open.Before(close) {
		return !eventTime.Before(open) && !eventTime.After(close)
	}
	if open.Equal(close) {
		return eventTime.Equal(open)
	}
	return !open.After(eventTime) || !close.Before(eventTime)
}

func CalculateProfit(pricePerHour int, seatingTime time.Duration) int {
	if seatingTime < 0 {
		seatingTime += 24 * time.Hour
	}
	inMinutes := int(seatingTime.Minutes())
	s := inMinutes / 60 * pricePerHour
	if inMinutes%60 > 0 {
		s += pricePerHour
	}
	return s
}
