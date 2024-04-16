package main

import (
	"net"

	"github.com/google/uuid"
)

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

func handleConnection(conn net.Conn, server *Server) {
	defer conn.Close()

	// Handle the connection. This could involve reading data from the
	// connection, updating the server state, etc.
}
