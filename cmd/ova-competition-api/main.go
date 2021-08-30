package main

import (
	"log"
	"ozonva/ova-competition-api/internal/api"
	"ozonva/ova-competition-api/internal/utils"
)

func main() {
	utils.WriteSomeEntries(10, "testfile.txt")

	if err := api.RunGrpcServer(); err != nil {
		log.Fatal(err)
	}
}
