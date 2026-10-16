# Go To-Do List Backend

A simple project I did to learn **Go**, **Gin**, and how to connect Go to a database (PostgreSQL via `pgx`).

Stack: Go 1.27.1 + Gin + `pgx/v5/pgxpool` + `godotenv` + PostgreSQL. Layered as `handler → service → repository`.

## Prerequisites

- Go 1.27+
- Docker (for PostgreSQL)

## 1. Start PostgreSQL

```powershell
docker run --name todo-pg `
  -e POSTGRES_PASSWORD=postgres `
  -e POSTGRES_DB=todos `
  -p 5432:5432 -d postgres:16
```

## 2. Create table + seed data

```powershell
docker exec -it todo-pg psql -U postgres -d todos
```

Then inside `psql` (`todos=#`):

```sql
CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  task TEXT NOT NULL,
  assignee TEXT NOT NULL
);

INSERT INTO todos (task, assignee) VALUES
  ('Vacuum the floor', 'Stan'),
  ('Wipe the floor', 'Stacy'),
  ('Cook Dinner', 'Mom'),
  ('Buy Grocery', 'Dad');
```

Verify: `SELECT * FROM todos;`. Check columns anytime with `\d todos`.

## 3. Configure env

```powershell
Copy-Item .env.example .env
```

Edit `.env`:

```
DATABASE_URL=postgresql://postgres:postgres@localhost:5432/todos?sslmode=disable
```

Format: `postgresql://USER:PASSWORD@HOST:PORT/DBNAME?sslmode=...`. Use `sslmode=disable` locally, `sslmode=require` on hosted DBs. The app auto-loads `.env` via `godotenv`; in prod just set a real `DATABASE_URL` env var (no `.env` file needed, it's gitignored).

## 4. Run

```powershell
go mod tidy
go run .
```

Server runs on `http://localhost:8080`.

## Endpoints

| Method | Path | Body | Success |
|---|---|---|---|
| GET | `/api/to-do-list` | — | 200 + list |
| POST | `/api/to-do-list` | `{"task":"...","assignee":"..."}` | 201 + created row |
| PUT | `/api/to-do-list/:id` | `{"task":"...","assignee":"..."}` | 200 + updated row, 404 if missing |
| DELETE | `/api/to-do-list/:id` | — | 204, 404 if missing |

```powershell
curl http://localhost:8080/api/to-do-list
curl -X POST http://localhost:8080/api/to-do-list -H "Content-Type: application/json" -d '{"task":"Mop floor","assignee":"Stan"}'
curl -X PUT http://localhost:8080/api/to-do-list/1 -H "Content-Type: application/json" -d '{"task":"Wipe floor twice","assignee":"Stacy"}'
curl -X DELETE http://localhost:8080/api/to-do-list/1 -i
```

## Notes

- One `pgxpool` is created in `main.go` and passed to handlers; request context (`c.Request.Context()`) is threaded to every query.
- All SQL uses `$1, $2` placeholders, so inputs are not vulnerable to SQL injection.
- `INSERT/UPDATE ... RETURNING id, task, assignee` writes and returns the row in one round-trip (also powers the 404 on update via `pgx.ErrNoRows`).
