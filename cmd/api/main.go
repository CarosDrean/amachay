package main

import (
	"log"

	"github.com/CarosDrean/amachay/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
