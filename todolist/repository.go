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

func insertToDoItem(ctx context.Context, pool *pgxpool.Pool, task, assignee string) (todoitem, error) {
	var item todoitem
	err := pool.QueryRow(ctx,
		`INSERT INTO todos (task, assignee) VALUES ($1, $2) RETURNING id, task, assignee`,
		task, assignee).Scan(&item.ID, &item.Task, &item.Assignee)
	return item, err
}

func updateToDoItem(ctx context.Context, pool *pgxpool.Pool, id int, task, assignee string) (todoitem, error) {
	var item todoitem
	err := pool.QueryRow(ctx,
		`UPDATE todos SET task = $1, assignee = $2 WHERE id = $3 RETURNING id, task, assignee`,
		task, assignee, id).Scan(&item.ID, &item.Task, &item.Assignee)
	return item, err
}

func deleteToDoItem(ctx context.Context, pool *pgxpool.Pool, id int) (int64, error) {
	res, err := pool.Exec(ctx, `DELETE FROM todos WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected(), nil
}
