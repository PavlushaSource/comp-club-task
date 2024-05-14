package internal

import (
	"fmt"
	"github.com/PavlushaSource/comp-club-task/internal/core/entity"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadNumberTables(data string, info *entity.ClubInfo) error {
	countTables, err := strconv.Atoi(data)
	if err != nil {
		return fmt.Errorf("error convert to int: %w", err)
	}
	if countTables <= 0 {
		return fmt.Errorf("the number must be positive")
	}

	info.NumberTables = countTables
	return nil
}

func ReadOpeningHours(data string, info *entity.ClubInfo) error {
	clubTime := strings.Split(data, " ")
	if len(clubTime) != 2 {
		return fmt.Errorf("failed parse opening hours. Format must be <15:04 20:04>")
	}
	openTime, err := time.Parse("15:04", clubTime[0])
	if err != nil {
		return fmt.Errorf("failed parse opening time: %w", err)
	}
	closeTime, err := time.Parse("15:04", clubTime[1])
	if err != nil {
		return fmt.Errorf("failed parse closing time: %w", err)
	}
	info.OpenTime = openTime
	info.CloseTime = closeTime

	return nil
}

func ReadTablePrice(data string, info *entity.ClubInfo) error {
	pricePerHour, err := strconv.Atoi(data)
	if err != nil {
		return fmt.Errorf("error convert to int: %w", err)
	}
	if pricePerHour <= 0 {
		return fmt.Errorf("the price must be positive")
	}

	info.PriceHour = pricePerHour
	return nil
}

func ReadFile(filepath string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error read input: %w", err)
	}
	if len(data) < 3 {
		return nil, fmt.Errorf("input is too short")
	}
	return strings.Split(string(data), "\n"), nil
}
