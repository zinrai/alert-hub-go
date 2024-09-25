package domain

import (
	"time"
)

type Alert struct {
	ID         int       `json:"id"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	Identifier string    `json:"identifier"`
	Urgency    string    `json:"urgency"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
