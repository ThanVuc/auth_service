package main

import (
	"auth_service/cmd"
	"log"
)

func main() {
	log.Println("gRPC servers are running...\n")
	cmd.RunGRPCServer()
	// run the grpc server
}
