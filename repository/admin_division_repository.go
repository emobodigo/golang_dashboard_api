package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/emobodigo/golang_dashboard_api/exception"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/domain"
)

type AdminDivisionRepository struct {
	db *sql.DB
}

func NewAdminDivisionRepository(passedDB *sql.DB) IAdminDivisionRepository {
	return &AdminDivisionRepository{
		db: passedDB,
	}
}

func (a *AdminDivisionRepository) Delete(ctx context.Context, id int) {
	SQL := "DELETE FROM `admin_division` WHERE `division_id` = ?"
	_, err := a.db.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (a *AdminDivisionRepository) FindAll(ctx context.Context) []domain.AdminDivision {
	SQL := "SELECT * FROM `admin_division`"
	rows, err := a.db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var adminDivisions []domain.AdminDivision
	for rows.Next() {
		adminDivision := domain.AdminDivision{}
		err := rows.Scan(&adminDivision.DivisionId, &adminDivision.DivisionName)
		helper.PanicIfError(err)
		adminDivisions = append(adminDivisions, adminDivision)
	}
	return adminDivisions
}

func (a *AdminDivisionRepository) FindById(ctx context.Context, id int) (domain.AdminDivision, error) {
	SQL := "SELECT * FROM `admin_division` WHERE `division_id` = ?"
	rows, err := a.db.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)

	adminDivision := domain.AdminDivision{}
	if rows.Next() {
		err := rows.Scan(&adminDivision.DivisionId, &adminDivision.DivisionName)
		helper.PanicIfError(err)
		return adminDivision, nil
	} else {
		return adminDivision, errors.New("admin division is not found")
	}
}

func (a *AdminDivisionRepository) Save(ctx context.Context, adminDivision domain.AdminDivision) domain.AdminDivision {
	duplicate := helper.CheckDuplicate(ctx, a.db, "admin_division", "division_name", adminDivision.DivisionName)
	if duplicate {
		err := errors.New("duplicate division name")
		panic(exception.NewConflictError(err.Error()))
	}
	SQL := "INSERT INTO `admin_division` (`division_name`) VALUES (?)"
	result, err := a.db.ExecContext(ctx, SQL, adminDivision.DivisionName)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	adminDivision.DivisionId = int(id)
	return adminDivision
}

func (a *AdminDivisionRepository) Update(ctx context.Context, adminDivision domain.AdminDivision) domain.AdminDivision {
	SQL := "UPDATE admin_division SET division_name = ? WHERE division_id = ?"
	_, err := a.db.ExecContext(ctx, SQL, adminDivision.DivisionName, adminDivision.DivisionId)
	helper.PanicIfError(err)
	return adminDivision
}
