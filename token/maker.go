package token

import (
	"time"
)

// Maker is an interface for managing tokens
type Maker interface {
	// Creates a new token for specific username and duration
	CreateToken(username string, duration time.Duration) (string, error)

	// Check if the token is valid or not
	VerifyToken(token string) (*Payload, error)
}
