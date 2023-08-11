package repository

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/model/domain"
)

type IAdminDivisionRepository interface {
	Save(ctx context.Context, adminDivision domain.AdminDivision) domain.AdminDivision
	Update(ctx context.Context, adminDivision domain.AdminDivision) domain.AdminDivision
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) (domain.AdminDivision, error)
	FindAll(ctx context.Context) []domain.AdminDivision
}
