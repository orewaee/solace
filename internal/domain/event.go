package domain

import "time"

type Event struct {
	Id        Id        `json:"id"`
	UserId    Id        `json:"user_id"`
	Text      string    `json:"text"`
	StartAt   time.Time `json:"start_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
