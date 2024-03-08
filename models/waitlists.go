package models

import "time"

type Waitlists struct {
	ID        uint   `json:"id"`
	Email     string `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
