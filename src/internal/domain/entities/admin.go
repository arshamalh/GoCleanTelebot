package entities

import (
	"time"
)

// Admin entity
type Admin struct {
	ID        string
	TID       string
	Active    bool
	CreatedAt time.Time
}
