package main

type GameMode int

const (
	Race GameMode = iota
	Coop
	CoopRace
)

type Game struct {
	// Base
	ID        int
	Name      string
	Health    int
	GameMode  GameMode
	Inventory Inventory
	ItemBox   ItemBox

	// Challenges
	SharedItemBox bool
	SharedHealth  bool
}

func InitializeGame(server *Server, channelUUID string, game string, gameMode GameMode) {
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
				GameMode:  gameMode,
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
