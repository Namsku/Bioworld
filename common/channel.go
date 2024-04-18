package common

import "github.com/google/uuid"

type Channel struct {
	// Private option
	Private  bool
	Password string

	// Default values
	UUID    string
	Players map[string]Player
	Game    Game
	Size    int
}

func CreateChannel(server *Server, password string, size int) {
	channel := Channel{
		UUID:     uuid.New().String(),
		Private:  password != "",
		Password: password,
		Players:  map[string]Player{},
		Game:     Game{},
		Size:     size,
	}

	Debug.Printf("Created channel with UUID %s\n", channel.UUID)
	server.Channels[channel.UUID] = channel
}

func DeleteChannel(server *Server, channelUUID string) {
	delete(server.Channels, channelUUID)
	Debug.Printf("Deleted channel with UUID %s\n", channelUUID)
}

func JoinChannel(server *Server, channelUUID string, password string, player Player) bool {
	channel, exists := server.Channels[channelUUID]
	if !exists {
		Debug.Printf("Channel with UUID %s does not exist\n", channelUUID)
		return false
	}

	if len(channel.Players) >= channel.Size {
		Debug.Printf("Channel with UUID %s is full\n", channelUUID)
		return false
	}

	if password != channel.Password {
		Debug.Printf("Incorrect password for channel with UUID %s\n", channelUUID)
	}

	channel.Players[player.UUID] = player
	server.Channels[channelUUID] = channel

	Debug.Printf("Player with UUID %s joined channel with UUID %s\n", player.UUID, channelUUID)
	return true
}

func LeaveChannel(server *Server, channelUUID string, playerUUID string) {
	channel, exists := server.Channels[channelUUID]
	if exists {
		delete(channel.Players, playerUUID)
		server.Channels[channelUUID] = channel
		Debug.Printf("Player with UUID %s left channel with UUID %s\n", playerUUID, channelUUID)
	}
}
