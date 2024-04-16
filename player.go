package main

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
