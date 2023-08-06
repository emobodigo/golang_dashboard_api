package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/domain"
)

type AdminDivisionRepository struct {
}

func (a *AdminDivisionRepository) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM `admin_division` WHERE `division_id` = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (a *AdminDivisionRepository) FindAll(ctx context.Context, tx *sql.Tx) []domain.AdminDivision {
	SQL := "SELECT * FROM `admin_division`"
	rows, err := tx.QueryContext(ctx, SQL)
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

func (a *AdminDivisionRepository) FindById(ctx context.Context, tx *sql.Tx, id int) (domain.AdminDivision, error) {
	SQL := "SELECT * FROM `admin_division` WHERE `division_id` = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
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

func (a *AdminDivisionRepository) Save(ctx context.Context, tx *sql.Tx, adminDivision domain.AdminDivision) domain.AdminDivision {
	SQL := "INSERT INTO `admin_division` (`division_name`) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, adminDivision)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	adminDivision.DivisionId = int(id)
	return adminDivision
}

func (a *AdminDivisionRepository) Update(ctx context.Context, tx *sql.Tx, adminDivision domain.AdminDivision) domain.AdminDivision {
	SQL := "UPDATE `admin_division` SET `division_name` = ? WHERE `division_id` = ?"
	_, err := tx.ExecContext(ctx, SQL, adminDivision.DivisionName, adminDivision.DivisionId)
	helper.PanicIfError(err)
	return adminDivision
}
