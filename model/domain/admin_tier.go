package domain

import "github.com/emobodigo/golang_dashboard_api/util"

type AdminTier struct {
	AdminTierId   util.StringInt
	AdminLevel    util.StringInt
	DivisionId    util.StringInt
	LevelTitle    string
	Fulltime      util.StringInt
	AdminDivision AdminDivision
}
