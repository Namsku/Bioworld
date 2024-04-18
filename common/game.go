package common

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

	// Co-Op Challenges
	SharedItemBox bool
	SharedHealth  bool

	// Classic Challenges
	NoDamage  bool
	NoItemBox bool
	NoSave    bool

	// Advanced Challenges
	RoomPenalty       bool
	EnemyKillingBonus bool
	RestTimer         bool
}

func InitializeGame(server *Server, channelUUID string, game string, gameMode GameMode) {
	var health int

	if game == "Biohazard" {
		health = 140
	}

	channel, exists := server.Channels[channelUUID]
	if exists {
		channel.Game = Game{
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
		server.Channels[channelUUID] = channel
	}
}
