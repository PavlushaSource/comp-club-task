package utils

import (
	"fmt"
	"github.com/PavlushaSource/comp-club-task/internal"
	"github.com/PavlushaSource/comp-club-task/internal/core/entity"
)

func ParseInfo(data []string) (*entity.ClubInfo, error) {
	var info entity.ClubInfo

	if err := internal.ReadNumberTables(data[0], &info); err != nil {
		return nil, fmt.Errorf("error number of tables: %w", err)
	}

	if err := internal.ReadOpeningHours(data[1], &info); err != nil {
		return nil, fmt.Errorf("error read opening hours: %w", err)
	}

	if err := internal.ReadTablePrice(data[2], &info); err != nil {
		return nil, fmt.Errorf("error read table price: %w", err)
	}

	return &info, nil
}
