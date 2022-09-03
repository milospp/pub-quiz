package config

import (
	"errors"
	"time"
)

type AuthClaim struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstname`
	Lastname  string    `json:"lastname"`
	Team      string    `json:"team"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func (claim *AuthClaim) Valid() error {
	if time.Now().After(claim.ExpiredAt) {
		return errors.New("token has expired")
	}
	return nil
}
