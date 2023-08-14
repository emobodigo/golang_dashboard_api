package services

import (
	"context"

	"github.com/emobodigo/golang_dashboard_api/exception"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/model/payload"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/go-playground/validator/v10"
)

type AdminTierService struct {
	AdminTierRepository repository.IAdminTierRepository
	Validate            *validator.Validate
}

func NewAdminTierService(adminTierRepository repository.IAdminTierRepository, validate *validator.Validate) IAdminTierService {
	return &AdminTierService{
		AdminTierRepository: adminTierRepository,
		Validate:            validate,
	}
}

func (a *AdminTierService) Create(ctx context.Context, request payload.AdminTierCreateRequest) payload.AdminTierResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	adminTier := domain.AdminTier{
		AdminLevel: request.AdminLevel,
		DivisionId: request.DivisionId,
		LevelTitle: request.LevelTitle,
		Fulltime:   request.Fulltime,
	}
	adminTier = a.AdminTierRepository.Save(ctx, adminTier)
	return helper.ToAdminTierResponse(adminTier)
}

func (a *AdminTierService) FindAllPaged(ctx context.Context, request payload.AdminTierPagedRequest) []payload.AdminTierResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	adminTiers := a.AdminTierRepository.FindAllPaged(ctx, request.Sort, request.Page, 15)
	return helper.ToAdminTierResponses(adminTiers)
}

func (a *AdminTierService) FindById(ctx context.Context, id int) payload.AdminTierResponse {
	adminTier, err := a.AdminTierRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToAdminTierResponse(adminTier)
}

func (a *AdminTierService) Update(ctx context.Context, request payload.AdminTierUpdateRequest) payload.AdminTierResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	adminTier, err := a.AdminTierRepository.FindById(ctx, request.AdminTierId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	adminTier.AdminLevel = request.AdminLevel
	adminTier.DivisionId = request.DivisionId
	adminTier.LevelTitle = request.LevelTitle
	adminTier.Fulltime = request.Fulltime

	adminTier = a.AdminTierRepository.Update(ctx, adminTier)
	return helper.ToAdminTierResponse(adminTier)
}
