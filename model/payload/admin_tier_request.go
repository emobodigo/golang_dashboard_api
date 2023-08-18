package payload

import "github.com/emobodigo/golang_dashboard_api/util"

type AdminTierCreateRequest struct {
	AdminLevel util.StringInt `json:"admin_level" validate:"required"`
	DivisionId util.StringInt `json:"division_id" validate:"required"`
	LevelTitle string         `json:"level_title" validate:"required"`
	Fulltime   util.StringInt `json:"fulltime" validate:"required"`
}

type AdminTierUpdateRequest struct {
	AdminTierId util.StringInt `json:"admin_tier_id" validate:"required"`
	AdminLevel  util.StringInt `json:"admin_level" validate:"required"`
	DivisionId  util.StringInt `json:"division_id" validate:"required"`
	LevelTitle  string         `json:"level_title" validate:"required"`
	Fulltime    util.StringInt `json:"fulltime" validate:"required"`
}

type AdminTierPagedRequest struct {
	Page util.StringInt `json:"page" validate:"required,min=1"`
	Sort string         `json:"sort"`
}
