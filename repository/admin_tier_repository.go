package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/emobodigo/golang_dashboard_api/exception"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/util"
)

type AdminTierRepository struct {
	db *sql.DB
}

func NewAdminTierRepository(db *sql.DB) IAdminTierRepository {
	return &AdminTierRepository{
		db: db,
	}
}

func (a *AdminTierRepository) FindAll(ctx context.Context) []domain.AdminTier {
	SQL := `SELECT at.*, ad.division_name
		FROM admin_tier at
		JOIN admin_division ad ON ad.division_id = at.division_id
		WHERE 1
	`
	rows, err := a.db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var adminTiers []domain.AdminTier
	for rows.Next() {
		adminTier := domain.AdminTier{}
		err := rows.Scan(&adminTier.AdminTierId, &adminTier.AdminLevel, &adminTier.DivisionId, &adminTier.LevelTitle, &adminTier.Fulltime, &adminTier.AdminDivision.DivisionName)
		helper.PanicIfError(err)
		adminTiers = append(adminTiers, adminTier)
	}
	return adminTiers
}

func (a *AdminTierRepository) FindById(ctx context.Context, id util.StringInt) (domain.AdminTier, error) {
	SQL := `SELECT at.*, ad.division_name
		FROM admin_tier at
		JOIN admin_division ad ON ad.division_id = at.division_id
		WHERE at.admin_tier_id = ?	
	`
	rows := a.db.QueryRowContext(ctx, SQL, id)
	adminTier := domain.AdminTier{}
	err := rows.Scan(&adminTier.AdminTierId, &adminTier.AdminLevel, &adminTier.DivisionId, &adminTier.LevelTitle, &adminTier.Fulltime, &adminTier.AdminDivision.DivisionName)
	if err != nil {
		return adminTier, errors.New("admin tier is not found")
	}
	return adminTier, nil
}

func (a *AdminTierRepository) Save(ctx context.Context, adminTier domain.AdminTier) domain.AdminTier {
	duplicate := helper.CheckDuplicate(ctx, a.db, "admin_tier", "level_title", adminTier.AdminLevel)
	if duplicate {
		err := errors.New("duplicate level title")
		panic(exception.NewConflictError(err.Error()))
	}
	SQL := `INSERT INTO admin_tier 
		(admin_level, division_id, level_title, fulltime) VALUES
		(?, ?, ?, ?)
	`
	result, err := a.db.ExecContext(ctx, SQL, adminTier.AdminLevel, adminTier.DivisionId, adminTier.LevelTitle, adminTier.Fulltime)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	adminTier.AdminTierId = util.StringInt(id)
	adminTier.Fulltime = 1

	createdAdminTier, _ := a.FindById(ctx, util.StringInt(id))
	adminTier.AdminDivision.DivisionName = createdAdminTier.AdminDivision.DivisionName
	return adminTier
}

func (a *AdminTierRepository) Update(ctx context.Context, adminTier domain.AdminTier) domain.AdminTier {
	SQL := `UPDATE admin_tier 
		SET admin_level = ?,
		division_id = ?,
		level_title = ?,
		fulltime = ?
		WHERE admin_tier_id = ?
	`
	_, err := a.db.ExecContext(ctx, SQL, adminTier.AdminLevel, adminTier.DivisionId, adminTier.LevelTitle, adminTier.Fulltime, adminTier.AdminTierId)
	helper.PanicIfError(err)
	return adminTier
}

func (a *AdminTierRepository) FindAllPaged(ctx context.Context, sort string, page util.StringInt, itemPerPage util.StringInt) []domain.AdminTier {
	start := (int(page) - 1) * int(itemPerPage)
	sortDir := "DESC"
	if strings.Contains(sort, "|") {
		sortDir = strings.Split(sort, "|")[1]
	}
	var sortSQL string
	switch sort {
	case "name":
		sortSQL = fmt.Sprintf("ORDER BY ad.division_name %v", sortDir)
	case "level":
		sortSQL = fmt.Sprintf("ORDER BY at.admin_level %v", sortDir)
	case "title":
		sortSQL = fmt.Sprintf("ORDER BY at.level_title %v", sortDir)
	default:
		sortSQL = fmt.Sprintf("ORDER BY at.admin_tier_id %v", sortDir)
	}

	SQL := fmt.Sprintf(`SELECT at.*, ad.division_name
		FROM admin_tier at
		JOIN admin_division ad ON ad.division_id = at.division_id
		WHERE 1
		%v
		LIMIT %v, %v
	`, sortSQL, start, itemPerPage)
	rows, err := a.db.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var adminTiers []domain.AdminTier
	for rows.Next() {
		adminTier := domain.AdminTier{}
		err := rows.Scan(&adminTier.AdminTierId, &adminTier.AdminLevel, &adminTier.DivisionId, &adminTier.LevelTitle, &adminTier.Fulltime, &adminTier.AdminDivision.DivisionName)
		helper.PanicIfError(err)
		adminTiers = append(adminTiers, adminTier)
	}
	return adminTiers
}
