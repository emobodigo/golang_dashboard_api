package repository

import (
	"context"
	"database/sql"

	"github.com/emobodigo/golang_dashboard_api/model/domain"
)

type IAdminDivisionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, adminDivision domain.AdminDivision) domain.AdminDivision
	Update(ctx context.Context, tx *sql.Tx, adminDivision domain.AdminDivision) domain.AdminDivision
	Delete(ctx context.Context, tx *sql.Tx, id int)
	FindById(ctx context.Context, tx *sql.Tx, id int) (domain.AdminDivision, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.AdminDivision
}
