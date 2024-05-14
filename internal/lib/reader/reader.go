package reader

import (
	"fmt"
	"os"
	"strings"
)

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
