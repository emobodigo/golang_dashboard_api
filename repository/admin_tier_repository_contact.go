package repository

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/model/domain"
)

type IAdminTierRepository interface {
	Save(ctx context.Context, adminTier domain.AdminTier) domain.AdminTier
	Update(ctx context.Context, adminTier domain.AdminTier) domain.AdminTier
	FindById(ctx context.Context, id int) (domain.AdminTier, error)
	FindAll(ctx context.Context) []domain.AdminTier
	FindAllPaged(ctx context.Context, sort string, page int, itemPerPage int) []domain.AdminTier
}
