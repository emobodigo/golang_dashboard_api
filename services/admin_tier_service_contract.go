package services

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/model/payload"
)

type IAdminTierService interface {
	Create(ctx context.Context, request payload.AdminTierCreateRequest) payload.AdminTierResponse
	Update(ctx context.Context, request payload.AdminTierUpdateRequest) payload.AdminTierResponse
	FindById(ctx context.Context, id int) payload.AdminTierResponse
	FindAllPaged(ctx context.Context, request payload.AdminTierPagedRequest) []payload.AdminTierResponse
}
