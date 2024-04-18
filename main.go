package main

import (
	"fmt"
	"log"
	"net"
<<<<<<< HEAD

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
=======
)

func main() {
	// Initialize your server here
	server := Server{
		Port:     1998,
		Password: "",
		Channels: []Channel{},
>>>>>>> 7feef6c833cdcbd1427daa577833e32ab4cd4894
	}
	// ...

	// Start a TCP server
<<<<<<< HEAD
	common.Debug.Println("Starting TCP Server...")
=======
>>>>>>> 7feef6c833cdcbd1427daa577833e32ab4cd4894
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server.Port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

<<<<<<< HEAD
	// Accept incoming connections
	common.Debug.Println("Accepting Incoming Connections...")
=======
>>>>>>> 7feef6c833cdcbd1427daa577833e32ab4cd4894
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

<<<<<<< HEAD
		go common.HandleConnection(conn, &server)
=======
		go handleConnection(conn, &server)
>>>>>>> 7feef6c833cdcbd1427daa577833e32ab4cd4894
	}
}
