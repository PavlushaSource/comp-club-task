package internal

import (
	"flag"
	"fmt"
	"os"
)

type Args struct {
	OutputPath string
	InputPath  string
}

func ParseFlags() (*Args, error) {
	var outputPath, inputPath string
	flag.StringVar(&outputPath, "output", "", "The output path, if empty, is output to the console")
	flag.StringVar(&inputPath, "input", "", "The input path, there must be a path to the input data file")
	flag.Parse()
	return &Args{
		OutputPath: outputPath,
		InputPath:  inputPath,
	}, nil
}

func ValidateFlags(args *Args) error {
	if args.OutputPath != "" {
		return fmt.Errorf("output to file is not supported yet")
	}

	if _, err := os.Stat(args.InputPath); os.IsNotExist(err) {
		return fmt.Errorf("input path <%s> does not exist", args.InputPath)
	}

	return nil
}
