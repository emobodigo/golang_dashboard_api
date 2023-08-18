package payload

import "github.com/emobodigo/golang_dashboard_api/util"

type AdminTierResponse struct {
	AdminTierId  util.StringInt `json:"admin_tier_id"`
	AdminLevel   util.StringInt `json:"admin_level"`
	DivisionId   util.StringInt `json:"division_id"`
	LevelTitle   string         `json:"level_title"`
	Fulltime     util.StringInt `json:"fulltime"`
	DivisionName string         `json:"division_name"`
}
