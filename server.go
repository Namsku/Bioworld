package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
)

type Slot struct {
	ID       int
	Quantity int
	Mode     int
}

type Inventory struct {
	SelectedSlot int
	Slots        []Slot
}

type ItemBox struct {
	Slots []Slot
}

type KeyItemStatus int

const (
	NotSeen KeyItemStatus = iota
	Seen
	Taken
)

type KeyItem struct {
	ID     int
	Name   string
	Status KeyItemStatus
}

type Game struct {
	// Base
	ID        int
	Name      string
	Health    int
	Inventory Inventory
	ItemBox   ItemBox

	// Challenges
	SharedItemBox bool
	SharedHealth  bool
}

type Player struct {
	// Server
	UUID        string
	Name        string
	Avatar      string
	ChannelUUID string

	// Game
	Health       int
	Room         string
	Timer        int
	LastItemSeen int
	Inventory    Inventory
	ItemBox      ItemBox

	// Stats
	Hit    int
	Deaths int
	Reset  int
}

type Channel struct {
	UUID     string
	Password string
	Players  []Player
	Game     Game
}

type Server struct {
	Port     int
	Password string
	Channels []Channel
}

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

func createChannel(server *Server, password string) {
	channel := Channel{
		UUID:     uuid.New().String(),
		Password: password,
		Players:  []Player{},
		Game:     Game{},
	}

	server.Channels = append(server.Channels, channel)
}

func joinChannel(server *Server, channelUUID string, player Player) {
	for i, channel := range server.Channels {
		if channel.UUID == channelUUID {
			server.Channels[i].Players = append(server.Channels[i].Players, player)
			break
		}
	}
}

func leaveChannel(server *Server, channelUUID string, playerUUID string) {
	for i, channel := range server.Channels {
		if channel.UUID == channelUUID {
			for j, player := range channel.Players {
				if player.UUID == playerUUID {
					server.Channels[i].Players = append(channel.Players[:j], channel.Players[j+1:]...)
					break
				}
			}
			break
		}
	}
}

func InitializeGame(server *Server, channelUUID string, game string) {
	var health int // Corrected here

	if game == "Biohazard" {
		health = 140
	}

	for i, channel := range server.Channels {
		if channel.UUID == channelUUID {
			server.Channels[i].Game = Game{
				ID:        0,
				Name:      game,
				Health:    health,
				Inventory: Inventory{},
				ItemBox: ItemBox{
					Slots: []Slot{
						{ID: 8, Quantity: 15, Mode: 0},
						{ID: 8, Quantity: 15, Mode: 0},
					},
				},
			}
			break
		}
	}
}

func handleConnection(conn net.Conn, server *Server) {
	defer conn.Close()

	// Handle the connection. This could involve reading data from the
	// connection, updating the server state, etc.
}
