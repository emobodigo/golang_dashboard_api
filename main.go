package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/emobodigo/golang_dashboard_api/app"
	"github.com/emobodigo/golang_dashboard_api/controller"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/emobodigo/golang_dashboard_api/services"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)
	port := os.Getenv("PORT")
	db := app.NewDB()
	validate := validator.New()

	adminDivisionRepo := repository.NewAdminDivisionRepository()
	adminDivisionService := services.NewAdminService(adminDivisionRepo, db, validate)
	adminDivisionController := controller.NewAdminDivisionController(adminDivisionService)

	router := httprouter.New()
	router.GET("/api/admindivisions", adminDivisionController.FindAll)
	router.GET("/api/admindivisions/:id", adminDivisionController.FindById)
	router.POST("/api/admindivisions", adminDivisionController.Create)
	router.PUT("/api/admindivisions/:id", adminDivisionController.Update)
	router.DELETE("/api/admindivisions/:id", adminDivisionController.Delete)

	server := http.Server{
		Addr:    "localhost:5001",
		Handler: router,
	}
	fmt.Println("Server Running on Port " + port)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
