package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Initialize your server here
	server := Server{
		Port:     1998,
		Password: "",
		Channels: []Channel{},
	}
	// ...

	// Start a TCP server
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn, &server)
	}
}
