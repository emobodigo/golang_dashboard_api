package services

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/model/payload"
)

type IAdminDivisionService interface {
	Create(ctx context.Context, request payload.AdminDivisionCreateRequest) payload.AdminDivisionResponse
	Update(ctx context.Context, request payload.AdminDivisionUpdateRequest) payload.AdminDivisionResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) payload.AdminDivisionResponse
	FindAll(ctx context.Context) []payload.AdminDivisionResponse
}
