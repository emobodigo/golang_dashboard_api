package services

import (
	"database/sql"

	"context"

	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/model/payload"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/go-playground/validator/v10"
)

type AdminDivisionService struct {
	AdminDivisionRepository repository.IAdminDivisionRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewAdminService(adminDivisionRepository repository.IAdminDivisionRepository, db *sql.DB, validate *validator.Validate) IAdminDivisionService {
	return &AdminDivisionService{
		AdminDivisionRepository: adminDivisionRepository,
		DB:                      db,
		Validate:                validate,
	}
}

func (a *AdminDivisionService) Create(ctx context.Context, request payload.AdminDivisionCreateRequest) payload.AdminDivisionResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := a.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	adminDivision := domain.AdminDivision{
		DivisionName: request.DivisionName,
	}
	adminDivision = a.AdminDivisionRepository.Save(ctx, tx, adminDivision)

	return helper.ToAdminDivisionResponse(adminDivision)
}

func (a *AdminDivisionService) Delete(ctx context.Context, id int) {
	tx, err := a.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = a.AdminDivisionRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	a.AdminDivisionRepository.Delete(ctx, tx, id)
}

func (a *AdminDivisionService) FindAll(ctx context.Context) []payload.AdminDivisionResponse {
	tx, err := a.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	adminDivisions := a.AdminDivisionRepository.FindAll(ctx, tx)

	return helper.ToAdminDivisionResponses(adminDivisions)
}

func (a *AdminDivisionService) FindById(ctx context.Context, id int) payload.AdminDivisionResponse {
	tx, err := a.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	adminDivision, err := a.AdminDivisionRepository.FindById(ctx, tx, id)
	helper.PanicIfError(err)

	return helper.ToAdminDivisionResponse(adminDivision)
}

func (a *AdminDivisionService) Update(ctx context.Context, request payload.AdminDivisionUpdateRequest) payload.AdminDivisionResponse {
	err := a.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := a.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	adminDivision, err := a.AdminDivisionRepository.FindById(ctx, tx, request.DivisionId)
	helper.PanicIfError(err)

	adminDivision.DivisionName = request.DivisionName

	adminDivision = a.AdminDivisionRepository.Update(ctx, tx, adminDivision)

	return helper.ToAdminDivisionResponse(adminDivision)
}
