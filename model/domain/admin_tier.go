package domain

type AdminTier struct {
	AdminTierId   int
	AdminLevel    int
	DivisionId    int
	LevelTitle    string
	Fulltime      int
	AdminDivision AdminDivision
}
