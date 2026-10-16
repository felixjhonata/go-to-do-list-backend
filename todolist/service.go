package todolist

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type todoitem struct {
	ID       int    `json:"id"`
	Task     string `json:"task"`
	Assignee string `json:"assignee"`
}

func listToDoItems(ctx context.Context, pool *pgxpool.Pool) ([]todoitem, error) {
	return getToDoItems(ctx, pool)
}