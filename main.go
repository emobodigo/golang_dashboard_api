package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/emobodigo/golang_dashboard_api/app"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/middleware"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helper.PanicIfError(err)
	port := os.Getenv("PORT")

	db := app.NewDB()
	validate := validator.New()

	authRouter := app.NewAuthRouter(db, validate)

	authHandler := middleware.NewAuthMiddleware(authRouter)

	server := http.Server{
		Addr:    "localhost:5001",
		Handler: authHandler,
	}
	fmt.Println("Server Running on Port " + port)
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
