// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: admin_division.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createAdminDivision = `-- name: CreateAdminDivision :execresult
INSERT INTO admin_division (
  division_name
) VALUES (
  ?
)
`

func (q *Queries) CreateAdminDivision(ctx context.Context, divisionName string) (sql.Result, error) {
	return q.db.ExecContext(ctx, createAdminDivision, divisionName)
}

const getAdminDivision = `-- name: GetAdminDivision :one
SELECT division_id, division_name FROM admin_division
WHERE division_id = ?
`

func (q *Queries) GetAdminDivision(ctx context.Context, divisionID int32) (AdminDivision, error) {
	row := q.db.QueryRowContext(ctx, getAdminDivision, divisionID)
	var i AdminDivision
	err := row.Scan(&i.DivisionID, &i.DivisionName)
	return i, err
}

const listAdminDivision = `-- name: ListAdminDivision :many
SELECT division_id, division_name FROM admin_division
`

func (q *Queries) ListAdminDivision(ctx context.Context) ([]AdminDivision, error) {
	rows, err := q.db.QueryContext(ctx, listAdminDivision)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []AdminDivision
	for rows.Next() {
		var i AdminDivision
		if err := rows.Scan(&i.DivisionID, &i.DivisionName); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAdminDivision = `-- name: UpdateAdminDivision :execresult
UPDATE admin_division 
SET division_name = ?
WHERE division_id = ?
`

type UpdateAdminDivisionParams struct {
	DivisionName string `json:"division_name"`
	DivisionID   int32  `json:"division_id"`
}

func (q *Queries) UpdateAdminDivision(ctx context.Context, arg UpdateAdminDivisionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, updateAdminDivision, arg.DivisionName, arg.DivisionID)
}
