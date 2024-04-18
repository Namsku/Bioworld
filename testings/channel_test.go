package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/namsku/bioworld/common"
)

func TestCreateChannelStress(t *testing.T) {
	// Initialize your logger here
	debugMode := true // Set this to true to enable debug logging
	common.InitLogger(debugMode)

	server := common.Server{
		Port:     1998,
		Password: "",
		Channels: map[string]common.Channel{}, // Change this line
	}

	// Create a large number of channels
	numChannels := 10000
	for i := 0; i < numChannels; i++ {
		common.CreateChannel(&server, "password", 4)
	}

	// Check if the correct number of channels was created
	if len(server.Channels) != numChannels {
		t.Errorf("Expected %d channels, got %d", numChannels, len(server.Channels))
	}
}

func TestServerStress(t *testing.T) {
	// Initialize your logger here
	debugMode := true // Set this to true to enable debug logging
	common.InitLogger(debugMode)

	server := common.Server{
		Port:     1998,
		Password: "",
		Channels: map[string]common.Channel{}, // Change this line
	}

	// Create a large number of channels
	numChannels := 10
	for i := 0; i < numChannels; i++ {
		common.CreateChannel(&server, "password", 4)
	}

	// Check if the correct number of channels was created
	if len(server.Channels) != numChannels {
		t.Errorf("Expected %d channels, got %d", numChannels, len(server.Channels))
	}

	numPlayers := 100
	numPlayersWithChannel := 0
	for i := 0; i < numPlayers; i++ {
		player := common.Player{
			UUID: uuid.New().String(),
			Name: "Player",
		}
		for uuid := range server.Channels {
			if common.JoinChannel(&server, uuid, "password", player) == true {
				numPlayersWithChannel++
				break
			}
		}
	}

	// Check if the correct number of players joined a channel
	if numChannels*4 != numPlayersWithChannel {
		t.Errorf("Expected %d players to join a channel, got %d", numPlayers-numChannels*4, numPlayersWithChannel)
	}

	// Delete all channels
	for uuid := range server.Channels {
		common.DeleteChannel(&server, uuid)
	}

	if len(server.Channels) != 0 {
		t.Errorf("Expected 0 channels, got %d", len(server.Channels))
	}
}
