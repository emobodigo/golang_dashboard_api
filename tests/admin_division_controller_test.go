package tests

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/emobodigo/golang_dashboard_api/app"
	"github.com/emobodigo/golang_dashboard_api/controller"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/middleware"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/emobodigo/golang_dashboard_api/services"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/dashboard_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	adminDivisionRepo := repository.NewAdminDivisionRepository()
	adminDivisionService := services.NewAdminService(adminDivisionRepo, db, validate)
	adminDivisionController := controller.NewAdminDivisionController(adminDivisionService)

	authRouter := app.NewAuthRouter(adminDivisionController)
	authHandler := middleware.NewAuthMiddleware(authRouter)
	return authHandler
}

func TestCreateAdminDivisionSuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"division_name": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5001/api/admindivisions", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	request.Header.Add("Authorization", "Bearer xaas")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, responseBody["code"])
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["division_name"])
}

func TestCreateAdminDivisionFailed(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"division_name": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5001/api/admindivisions", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	request.Header.Add("Authorization", "Bearer xaas")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 409, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(responseBody)

	assert.Equal(t, 409, int(responseBody["code"].(float64)))
	assert.Equal(t, "Conflict Duplicate", responseBody["status"])
	assert.Equal(t, "duplicate division name", responseBody["data"])
}

func TestUpdateAdminDivisionSuccess(t *testing.T) {

}

func TestUpdateAdminDivisionFailed(t *testing.T) {

}

func TestDeleteAdminDivisionSuccess(t *testing.T) {

}

func TestDeleteAdminDivisionFailed(t *testing.T) {

}

func TestGetAdminDivisionSuccess(t *testing.T) {

}

func TestGetAdminDivisionFailed(t *testing.T) {

}

func TestGetAllAdminDivisionSuccess(t *testing.T) {

}

func TestGetAllAdminDivisionFailed(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
