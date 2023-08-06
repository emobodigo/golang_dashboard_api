package payload

type AdminDivisionUpdateRequest struct {
	DivisionId   int    `validate:"required"`
	DivisionName string `validate:"required" json:"division_name"`
}
