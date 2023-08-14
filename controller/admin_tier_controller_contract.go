package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IAdminTierController interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindAllPaged(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
