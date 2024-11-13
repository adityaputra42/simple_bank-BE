package token

import "time"

// Maker is an interface to managing token
type Maker interface {
	// CreateToken create a new token for a specific username and duration
	CreateToken(username string, userId int64, duration time.Duration) (string, *Payload, error)
	// VerifyToken check if the token is valid or
	VerifyToken(token string) (*Payload, error)
}
