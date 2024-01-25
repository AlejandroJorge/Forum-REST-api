package model

import "time"

type User struct {
	ID uint

	// For auth
	Email          string
	HashedPassword string

	// Meta
	RegistrationDate time.Time
}
