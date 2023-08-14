package app

import (
	"database/sql"

	"github.com/emobodigo/golang_dashboard_api/controller"
	"github.com/emobodigo/golang_dashboard_api/exception"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/emobodigo/golang_dashboard_api/services"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func NewAuthRouter(db *sql.DB, validate *validator.Validate) *httprouter.Router {
	router := httprouter.New()

	adminDivisionRepo := repository.NewAdminDivisionRepository(db)
	adminDivisionService := services.NewAdminDivisionService(adminDivisionRepo, validate)
	adminDivisionController := controller.NewAdminDivisionController(adminDivisionService)

	adminTierRepo := repository.NewAdminTierRepository(db)
	adminTierService := services.NewAdminTierService(adminTierRepo, validate)
	adminTierController := controller.NewAdminTierController(adminTierService)

	router.GET("/api/admindivisions", adminDivisionController.FindAll)
	router.GET("/api/admindivisions/:id", adminDivisionController.FindById)
	router.POST("/api/admindivisions", adminDivisionController.Create)
	router.PUT("/api/admindivisions/:id", adminDivisionController.Update)
	router.DELETE("/api/admindivisions/:id", adminDivisionController.Delete)

	router.GET("/api/admintiers", adminTierController.FindAllPaged)
	router.GET("/api/admintiers/:id", adminTierController.FindById)
	router.POST("/api/admintiers", adminTierController.Create)
	router.PUT("/api/admintiers/:id", adminTierController.Update)

	router.PanicHandler = exception.ErrorHandler

	return router
}
