package helper

import (
	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/model/payload"
)

func ToAdminDivisionResponse(adminDivision domain.AdminDivision) payload.AdminDivisionResponse {
	return payload.AdminDivisionResponse{
		DivisionId:   adminDivision.DivisionId,
		DivisionName: adminDivision.DivisionName,
	}
}

func ToAdminDivisionResponses(divisions []domain.AdminDivision) []payload.AdminDivisionResponse {
	var divisionResponses []payload.AdminDivisionResponse
	for _, division := range divisions {
		divisionResponses = append(divisionResponses, ToAdminDivisionResponse(division))
	}
	return divisionResponses
}
