-- name: GetTodo :one

SELECT * FROM todos WHERE id = ? LIMIT 1;

-- name: ListTodos :many

SELECT * FROM todos ORDER BY name;

-- name: CreateTodo :one

INSERT INTO todos ( name, complete ) VALUES ( ?, ? ) RETURNING *;

-- name: UpdateTodo :one

UPDATE todos set name = ?, complete = ? WHERE id = ? RETURNING *;

-- name: DeleteTodo :exec

DELETE FROM todos WHERE id = ?;