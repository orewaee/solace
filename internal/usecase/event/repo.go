package event

import (
	"context"

	"orewaee.dev/solace/internal/domain"
)

type Repo interface {
	Save(context.Context, *domain.Event) error
	FindById(context.Context, domain.Id)
	FindByUserId(context.Context, domain.Id)
	Delete(context.Context, domain.Id)
}
