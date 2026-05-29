package domain

import "time"

type Event struct {
	Id          Id        `json:"id"`
	UserId      Id        `json:"user_id"`
	Text        string    `json:"text"`
	ScheduledAt time.Time `json:"scheduled_at"`
	CreatedAt   time.Time `json:"updated_at"`
}
