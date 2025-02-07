package main

import (
	"log"

	"github.com/SigurdRiseth/CountryInfoService/server"
)

func main() {
	log.Println("Starting Country Info Service...")
	server.StartServer()
}
