package common

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
