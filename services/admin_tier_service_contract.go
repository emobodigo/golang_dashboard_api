package services

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/model/payload"
	"github.com/emobodigo/golang_dashboard_api/util"
)

type IAdminTierService interface {
	Create(ctx context.Context, request payload.AdminTierCreateRequest) payload.AdminTierResponse
	Update(ctx context.Context, request payload.AdminTierUpdateRequest) payload.AdminTierResponse
	FindById(ctx context.Context, id util.StringInt) payload.AdminTierResponse
	FindAllPaged(ctx context.Context, request payload.AdminTierPagedRequest) []payload.AdminTierResponse
}
