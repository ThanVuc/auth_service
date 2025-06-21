package main

import (
	"auth_service/internal/initialize"
	"log"
)

func main() {
	log.Println("gRPC servers are running...\n")
	initialize.Run()
	// run the grpc server
}
