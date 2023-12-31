// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package sqlc

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAdminDivision(ctx context.Context, divisionName string) (sql.Result, error)
	GetAdminDivision(ctx context.Context, divisionID int32) (AdminDivision, error)
	ListAdminDivision(ctx context.Context) ([]AdminDivision, error)
	UpdateAdminDivision(ctx context.Context, arg UpdateAdminDivisionParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
