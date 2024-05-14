package flag

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	InputPath string
}

func ParseFlags() (*Args, error) {
	var inputPath string
	flag.StringVar(&inputPath, "input", "", "The input path, there must be a path to the input data file")
	flag.Parse()
	if err := ValidatePath(inputPath); err != nil {
		return nil, fmt.Errorf("error validate path: %w", err)
	}
	return &Args{
		InputPath: inputPath,
	}, nil
}

func ValidatePath(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("input path <%s> does not exist", path)
	}
	return nil
}
