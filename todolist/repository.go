package todolist

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func getToDoItems(ctx context.Context, pool *pgxpool.Pool) ([]todoitem, error) {
	rows, err := pool.Query(ctx, `SELECT id, task, assignee FROM todos ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []todoitem{}
	for rows.Next() {
		var item todoitem
		if err := rows.Scan(&item.ID, &item.Task, &item.Assignee); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, rows.Err()
}