package controller

import (
	"net/http"
	"strconv"

	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/payload"
	"github.com/emobodigo/golang_dashboard_api/services"
	"github.com/julienschmidt/httprouter"
)

type AdminDivisionController struct {
	AdminDivisionService services.IAdminDivisionService
}

func NewAdminDivisionController(adminDivisionService services.IAdminDivisionService) IAdminDivisionController {
	return &AdminDivisionController{
		AdminDivisionService: adminDivisionService,
	}
}

func (a *AdminDivisionController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	divisionCreateRequest := payload.AdminDivisionCreateRequest{}
	helper.ReadFromRequestBody(r, &divisionCreateRequest)

	divisionResponse := a.AdminDivisionService.Create(r.Context(), divisionCreateRequest)
	apiResponse := payload.ApiResponse{
		Code:   201,
		Status: "OK",
		Data:   divisionResponse,
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminDivisionController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	divisionId := params.ByName("id")
	id, err := strconv.Atoi(divisionId)
	helper.PanicIfError(err)

	a.AdminDivisionService.Delete(r.Context(), id)
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminDivisionController) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	divisions := a.AdminDivisionService.FindAll(r.Context())
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   divisions,
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminDivisionController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	divisionId := params.ByName("id")
	id, err := strconv.Atoi(divisionId)
	helper.PanicIfError(err)

	division := a.AdminDivisionService.FindById(r.Context(), id)
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   division,
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminDivisionController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	divisionUpdateRequest := payload.AdminDivisionUpdateRequest{}
	helper.ReadFromRequestBody(r, &divisionUpdateRequest)

	divisionId := params.ByName("id")
	id, err := strconv.Atoi(divisionId)
	helper.PanicIfError(err)

	divisionUpdateRequest.DivisionId = id

	divisionResponse := a.AdminDivisionService.Update(r.Context(), divisionUpdateRequest)
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   divisionResponse,
	}
	helper.WriteToResponseBody(w, apiResponse)
}
