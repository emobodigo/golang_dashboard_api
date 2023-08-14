package payload

type AdminTierCreateRequest struct {
	AdminLevel int    `json:"admin_level" validate:"required"`
	DivisionId int    `json:"division_id" validate:"required"`
	LevelTitle string `json:"level_title" validate:"required"`
	Fulltime   int    `json:"fulltime" validate:"required, oneof=0 1"`
}

type AdminTierUpdateRequest struct {
	AdminTierId int    `json:"admin_tier_id" validate:"required"`
	AdminLevel  int    `json:"admin_level" validate:"required"`
	DivisionId  int    `json:"division_id" validate:"required"`
	LevelTitle  string `json:"level_title" validate:"required"`
	Fulltime    int    `json:"fulltime" validate:"required, oneof=0 1"`
}

type AdminTierPagedRequest struct {
	Page int    `json:"page" validate:"required, min=1"`
	Sort string `json:"sort"`
}
