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

func ToAdminTierResponse(tier domain.AdminTier) payload.AdminTierResponse {
	return payload.AdminTierResponse{
		AdminTierId:  tier.AdminTierId,
		AdminLevel:   tier.AdminLevel,
		DivisionId:   tier.DivisionId,
		LevelTitle:   tier.LevelTitle,
		DivisionName: tier.AdminDivision.DivisionName,
	}
}

func ToAdminTierResponses(adminTier []domain.AdminTier) []payload.AdminTierResponse {
	var tierResponses []payload.AdminTierResponse
	for _, tier := range adminTier {
		tierResponses = append(tierResponses, ToAdminTierResponse(tier))
	}
	if len(tierResponses) == 0 {
		return []payload.AdminTierResponse{}
	}
	return tierResponses
}
