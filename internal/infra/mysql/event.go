package mysql

import (
	"context"
	"database/sql"

	"orewaee.dev/solace/internal/db"
	"orewaee.dev/solace/internal/domain"
	"orewaee.dev/solace/internal/usecase/event"
)

type mysqlEventRepo struct {
	queries *db.Queries
}

func (r *mysqlEventRepo) Delete(context.Context, domain.Id) {
	panic("unimplemented")
}

func (r *mysqlEventRepo) FindById(context.Context, domain.Id) {
	panic("unimplemented")
}

func (r *mysqlEventRepo) FindByUserId(context.Context, domain.Id) {
	panic("unimplemented")
}

func (r *mysqlEventRepo) Save(ctx context.Context, event *domain.Event) error {
	arg := db.InsertEventParams{
		Id:     event.Id,
		UserId: event.UserId,
		Text: sql.NullString{
			String: event.Text,
			Valid:  len(event.Text) != 0,
		},
		StartAt:   event.StartAt,
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}

	_, err := r.queries.InsertEvent(ctx, arg)
	return err
}

func NewEventRepo(queries *db.Queries) event.Repo {
	return &mysqlEventRepo{
		queries: queries,
	}
}
