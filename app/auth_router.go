package app

import (
	"github.com/emobodigo/golang_dashboard_api/controller"
	"github.com/emobodigo/golang_dashboard_api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewAuthRouter(adminDivisionController controller.IAdminDivisionController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/admindivisions", adminDivisionController.FindAll)
	router.GET("/api/admindivisions/:id", adminDivisionController.FindById)
	router.POST("/api/admindivisions", adminDivisionController.Create)
	router.PUT("/api/admindivisions/:id", adminDivisionController.Update)
	router.DELETE("/api/admindivisions/:id", adminDivisionController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
