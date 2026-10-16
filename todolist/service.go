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

func createToDoItem(ctx context.Context, pool *pgxpool.Pool, task, assignee string) (todoitem, error) {
	return insertToDoItem(ctx, pool, task, assignee)
}

func editToDoItem(ctx context.Context, pool *pgxpool.Pool, id int, task, assignee string) (todoitem, error) {
	return updateToDoItem(ctx, pool, id, task, assignee)
}

func removeToDoItem(ctx context.Context, pool *pgxpool.Pool, id int) (int64, error) {
	return deleteToDoItem(ctx, pool, id)
}