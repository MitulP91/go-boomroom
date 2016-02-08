package "song"

import (
	"database/sql"
	"github.com/lib/pq"
	"db/connect"
)

type Song struct {
	// URL to the song content
	Url string

	// The website that this song was pulled from.
	Source string
}

// Persist the song to the PostgreSQL db
func (s *Song) persist() {
	// Check to see if we have an existing row for this song
	// in the DB based on the URL and Source of the song

}

// Constructor method
func NewSong() *Song {
	return &Song{}
}