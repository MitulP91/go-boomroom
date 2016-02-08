package room

import (
	"user"
)

type Room struct {
	// User input room name
	Name string
}

// Constructor method
func NewRoom() *Room {
	return &Room{}
}