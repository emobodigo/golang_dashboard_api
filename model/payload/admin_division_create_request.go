package payload

type AdminDivisionCreateRequest struct {
	DivisionName string `validate:"required" json:"division_name"`
}
