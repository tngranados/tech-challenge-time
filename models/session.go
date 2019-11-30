package models

import "time"

// Session represents a time tracking session in the system.
type Session struct {
	ID         int
	Name       string
	StartedAt  time.Time
	FinishedAt time.Time
}
