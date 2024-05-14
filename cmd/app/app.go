package main

import (
	"github.com/PavlushaSource/comp-club-task/internal/core/service"
	"github.com/PavlushaSource/comp-club-task/internal/lib/flag"
	"github.com/PavlushaSource/comp-club-task/internal/lib/parser"
	"github.com/PavlushaSource/comp-club-task/internal/lib/reader"
	"log"
)

func main() {
	args, err := flag.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	data, err := reader.ReadFile(args.InputPath)
	if err != nil {
		log.Fatal(err)
	}

	info, err := parser.ParseHeaderCompClub(data)
	if err != nil {
		log.Fatal(err)
	}

	err = service.StartClub(info, data[3:])
	if err != nil {
		log.Fatal(err)
	}
}
