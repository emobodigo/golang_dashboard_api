package repository

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/util"
)

type IAdminTierRepository interface {
	Save(ctx context.Context, adminTier domain.AdminTier) domain.AdminTier
	Update(ctx context.Context, adminTier domain.AdminTier) domain.AdminTier
	FindById(ctx context.Context, id util.StringInt) (domain.AdminTier, error)
	FindAll(ctx context.Context) []domain.AdminTier
	FindAllPaged(ctx context.Context, sort string, page util.StringInt, itemPerPage util.StringInt) []domain.AdminTier
}
