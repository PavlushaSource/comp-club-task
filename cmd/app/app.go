package main

import (
	"github.com/PavlushaSource/comp-club-task/internal"
	"github.com/PavlushaSource/comp-club-task/internal/core/service"
	"github.com/PavlushaSource/comp-club-task/internal/utils"
	"log"
)

func main() {
	args, err := internal.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	err = internal.ValidateFlags(args)
	if err != nil {
		log.Fatal(err)
	}

	data, err := internal.ReadFile(args.InputPath)
	if err != nil {
		log.Fatal(err)
	}
	info, err := utils.ParseInfo(data)
	if err != nil {
		log.Fatal(err)
	}

	err = service.StartClub(info, data[3:])
	if err != nil {
		log.Fatal(err)
	}
}
