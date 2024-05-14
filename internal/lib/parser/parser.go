package parser

import (
	"fmt"
	"github.com/PavlushaSource/comp-club-task/internal/core/entity"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseEvent(data string, info *entity.ClubInfo) (*entity.Event, error) {
	dataSplit := strings.Split(data, " ")
	if len(dataSplit) < 3 {
		return nil, fmt.Errorf("failed parse event, length must be 3 or larger, got %d", len(data))
	}
	timeEvent, err := time.Parse("15:04", dataSplit[0])
	if err != nil {
		return nil, fmt.Errorf("failed parse time event: %s", dataSplit[0])
	}

	eventID, err := strconv.Atoi(dataSplit[1])
	if err != nil {
		return nil, fmt.Errorf("failed parse event id: %s", dataSplit[1])
	}

	IsCorrectClientName := regexp.MustCompile(`^[a-z0-9_-]+$`).MatchString
	if !IsCorrectClientName(dataSplit[2]) {
		return nil, fmt.Errorf("failed parse client name: %s", dataSplit[2])
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
		if err != nil || tableNumber <= 0 || tableNumber > info.NumberTables {
			return nil, fmt.Errorf("uncorrect table number: %s", dataSplit[3])
		}
		event.TableNumber = tableNumber

	default:
		return nil, fmt.Errorf("uncorrect event id %d", eventID)
	}
	return event, nil
}
func ParseHeaderCompClub(data []string) (*entity.ClubInfo, error) {
	var info entity.ClubInfo

	if err := parseNumberTables(data[0], &info); err != nil {
		return nil, fmt.Errorf("error number of tables: %w", err)
	}

	if err := parseOpeningHours(data[1], &info); err != nil {
		return nil, fmt.Errorf("error read opening hours: %w", err)
	}

	if err := parseTablePrice(data[2], &info); err != nil {
		return nil, fmt.Errorf("error read table price: %w", err)
	}

	return &info, nil
}

func parseNumberTables(data string, info *entity.ClubInfo) error {
	countTables, err := strconv.Atoi(data)
	if err != nil {
		return fmt.Errorf("error convert to int: %s", data)
	}
	if countTables <= 0 {
		return fmt.Errorf("the number must be positive, got %d", countTables)
	}

	info.NumberTables = countTables
	return nil
}

func parseOpeningHours(data string, info *entity.ClubInfo) error {
	clubTime := strings.Split(data, " ")
	if len(clubTime) != 2 {
		return fmt.Errorf("failed parse opening hours. \nFormat must be <15:04 20:04>, got %s", data)
	}
	openTime, err := time.Parse("15:04", clubTime[0])
	if err != nil {
		return fmt.Errorf("failed parse opening time: %s", clubTime[0])
	}
	closeTime, err := time.Parse("15:04", clubTime[1])
	if err != nil {
		return fmt.Errorf("failed parse closing time: %s", clubTime[1])
	}
	info.OpenTime = openTime
	info.CloseTime = closeTime

	return nil
}

func parseTablePrice(data string, info *entity.ClubInfo) error {
	pricePerHour, err := strconv.Atoi(data)
	if err != nil {
		return fmt.Errorf("error convert to int: %s", data)
	}
	if pricePerHour <= 0 {
		return fmt.Errorf("the price must be positive, got %d", pricePerHour)
	}

	info.PriceHour = pricePerHour
	return nil
}
