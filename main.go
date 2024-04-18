package main

import (
	"fmt"
	"log"
	"net"

	"github.com/namsku/bioworld/common"
)

func main() {
	// Initialize your logger here
	debugMode := false // Set this to true to enable debug logging
	common.InitLogger(debugMode)

	common.Debug.Println("Creating Server Object...")
	server := common.Server{
		Port:     1998,
		Password: "",
		Channels: []common.Channel{},
	}
	// ...

	// Start a TCP server
	common.Debug.Println("Starting TCP Server...")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	// Accept incoming connections
	common.Debug.Println("Accepting Incoming Connections...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go common.HandleConnection(conn, &server)
	}
}
