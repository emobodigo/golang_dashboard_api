package controller

import (
	"net/http"
	"strconv"

	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/payload"
	"github.com/emobodigo/golang_dashboard_api/services"
	"github.com/emobodigo/golang_dashboard_api/util"
	"github.com/julienschmidt/httprouter"
)

type AdminTierController struct {
	AdminTierService services.IAdminTierService
}

func NewAdminTierController(adminTierService services.IAdminTierService) IAdminTierController {
	return &AdminTierController{
		AdminTierService: adminTierService,
	}
}

func (a *AdminTierController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	adminTierCreateRequest := payload.AdminTierCreateRequest{}
	helper.ReadFromRequestBody(r, &adminTierCreateRequest)

	tierResponse := a.AdminTierService.Create(r.Context(), adminTierCreateRequest)
	apiResponse := payload.ApiResponse{
		Code:   201,
		Status: "OK",
		Data:   tierResponse,
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminTierController) FindAllPaged(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	queryValues := r.URL.Query()
	pageString := queryValues.Get("page")
	sort := queryValues.Get("sort")
	if pageString == "" {
		pageString = "1"
	}
	page, err := strconv.Atoi(pageString)
	helper.PanicIfError(err)

	adminTierReq := payload.AdminTierPagedRequest{}
	adminTierReq.Page = util.StringInt(page)
	adminTierReq.Sort = sort

	tiers := a.AdminTierService.FindAllPaged(r.Context(), adminTierReq)
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   tiers,
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminTierController) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	adminTierId := params.ByName("id")
	id, err := strconv.Atoi(adminTierId)
	helper.PanicIfError(err)

	tier := a.AdminTierService.FindById(r.Context(), util.StringInt(id))
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   tier,
	}
	helper.WriteToResponseBody(w, apiResponse)
}

func (a *AdminTierController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	tierUpdateRequest := payload.AdminTierUpdateRequest{}
	helper.ReadFromRequestBody(r, &tierUpdateRequest)

	adminTierId := params.ByName("id")
	id, err := strconv.Atoi(adminTierId)
	helper.PanicIfError(err)

	tierUpdateRequest.AdminTierId = util.StringInt(id)
	tier := a.AdminTierService.Update(r.Context(), tierUpdateRequest)
	apiResponse := payload.ApiResponse{
		Code:   200,
		Status: "OK",
		Data:   tier,
	}
	helper.WriteToResponseBody(w, apiResponse)
}
