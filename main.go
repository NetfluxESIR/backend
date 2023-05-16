package main

import (
	"github.com/NetfluxESIR/backend/cmd"
	"log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
