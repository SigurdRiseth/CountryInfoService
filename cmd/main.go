package main

import (
	"github.com/SigurdRiseth/CountryInfoService/server"
	"log"
)

func main() {
	log.Println("Starting Country Info Service...")
	server.StartServer()
}
