package utils

import (
	"fmt"
	"github.com/PavlushaSource/comp-club-task/internal/core/entity"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseEvent(data string) (*entity.Event, error) {
	dataSplit := strings.Split(data, " ")
	if len(dataSplit) < 3 {
		return nil, fmt.Errorf("failed parse event, length must be 3 or larger")
	}
	timeEvent, err := time.Parse("15:04", dataSplit[0])
	if err != nil {
		return nil, fmt.Errorf("failed parse time event: %w", err)
	}

	eventID, err := strconv.Atoi(dataSplit[1])
	if err != nil {
		return nil, fmt.Errorf("failed parse event id: %w", err)
	}

	IsCorrectClientName := regexp.MustCompile(`^[a-z0-9_-]+$`).MatchString
	if !IsCorrectClientName(dataSplit[2]) {
		return nil, fmt.Errorf("failed parse client name: %w", err)
	}
	clientName := dataSplit[2]

	event := &entity.Event{
		Time:       timeEvent,
		ID:         eventID,
		ClientName: clientName,
	}

	switch {
	case eventID == entity.ClientWaiting || eventID == entity.ClientCome || eventID == entity.ClientLeft:
		if len(dataSplit) != 3 {
			return nil, fmt.Errorf("failed parse event: %s: input length must be 3", data)
		}
	case eventID == entity.ClientSeatTable:
		if len(dataSplit) != 4 {
			return nil, fmt.Errorf("failed parse event: %s: input length must be 4", data)
		}
		var tableNumber int
		tableNumber, err = strconv.Atoi(dataSplit[3])
		if err != nil || tableNumber <= 0 {
			return nil, fmt.Errorf("uncorrect table number: %w", err)
		}
		event.TableNumber = tableNumber

	default:
		return nil, fmt.Errorf("uncorrect event id %d", eventID)
	}
	return event, nil
}
