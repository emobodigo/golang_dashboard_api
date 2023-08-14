package payload

type AdminTierResponse struct {
	AdminTierId  int    `json:"admin_tier_id"`
	AdminLevel   int    `json:"admin_level"`
	DivisionId   int    `json:"division_id"`
	LevelTitle   string `json:"level_title"`
	Fulltime     int    `json:"fulltime"`
	DivisionName string `json:"division_name"`
}
