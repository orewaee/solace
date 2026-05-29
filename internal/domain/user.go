package domain

import "time"

type User struct {
	Id         Id         `json:"id"`
	TelegramId TelegramId `json:"telegram_id"`
	Name       string     `json:"name"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}
