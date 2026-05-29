package user

import (
	"context"

	"orewaee.dev/solace/internal/domain"
)

type User interface {
	Create(context.Context, *domain.User) error
	Update(context.Context, *User) error
	DeleteById(context.Context, domain.Id) error
	FindById(context.Context, domain.Id) (*User, error)
	FindByTelegramId(context.Context, domain.TelegramId) (*User, error)
}
