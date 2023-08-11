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

type AdminDivisionService struct {
	AdminDivisionRepository repository.IAdminDivisionRepository
	Validate                *validator.Validate
}

func NewAdminService(adminDivisionRepository repository.IAdminDivisionRepository, validate *validator.Validate) IAdminDivisionService {
	return &AdminDivisionService{
		AdminDivisionRepository: adminDivisionRepository,
		Validate:                validate,
	}
}

func (a *AdminDivisionService) Create(ctx context.Context, request payload.AdminDivisionCreateRequest) payload.AdminDivisionResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	adminDivision := domain.AdminDivision{
		DivisionName: request.DivisionName,
	}
	adminDivision = a.AdminDivisionRepository.Save(ctx, adminDivision)

	return helper.ToAdminDivisionResponse(adminDivision)
}

func (a *AdminDivisionService) Delete(ctx context.Context, id int) {
	_, err := a.AdminDivisionRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	a.AdminDivisionRepository.Delete(ctx, id)
}

func (a *AdminDivisionService) FindAll(ctx context.Context) []payload.AdminDivisionResponse {

	adminDivisions := a.AdminDivisionRepository.FindAll(ctx)

	return helper.ToAdminDivisionResponses(adminDivisions)
}

func (a *AdminDivisionService) FindById(ctx context.Context, id int) payload.AdminDivisionResponse {

	adminDivision, err := a.AdminDivisionRepository.FindById(ctx, id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToAdminDivisionResponse(adminDivision)
}

func (a *AdminDivisionService) Update(ctx context.Context, request payload.AdminDivisionUpdateRequest) payload.AdminDivisionResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	adminDivision, err := a.AdminDivisionRepository.FindById(ctx, request.DivisionId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	adminDivision.DivisionName = request.DivisionName

	adminDivision = a.AdminDivisionRepository.Update(ctx, adminDivision)

	return helper.ToAdminDivisionResponse(adminDivision)
}
