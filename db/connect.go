package "connect"

import (
	"database/sql"
	"github.com/lib/pq"
)

type PostgresConnector struct {
	// Connection username
	Username string

	// Connection password
	Password string

	// DB Name
	DbName string

	// DB host
	Host string
}

// Connect method on the PostgresConnector. Returns the
// PostgreSQL connection struct
func (pc *PostgresConnector) connect() conn, err {
	db, err := sql.Open("postgres", "user=")

	if err != nil {
		// Need to figure out how Go errors work :P
	}

	return conn, nil
}

// Constructor method
func NewPostgresConnector() *PostgresConnector {
	return &PostgresConnector{}
}