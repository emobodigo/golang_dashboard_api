package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/emobodigo/golang_dashboard_api/app"
	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/middleware"
	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/dashboard_api_test")
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/dashboard_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func clearDivisionData(db *sql.DB) {
	_, err := db.Exec("DELETE FROM admin_division WHERE division_id > 3")
	if err != nil {
		panic(err)
	}
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	authRouter := app.NewAuthRouter(db, validate)
	authHandler := middleware.NewAuthMiddleware(authRouter)
	return authHandler
}

func setupHeaderWithAuth(request *http.Request) {
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")
	request.Header.Add("Authorization", "Bearer xaas")
}

func TestCreateAdminDivisionSuccess(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	clearDivisionData(db)

	requestBody := strings.NewReader(`{"division_name": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5001/api/admindivisions", requestBody)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["division_name"])
}

func TestCreateAdminDivisionFailed(t *testing.T) {
	db := setupTestDB()
	router := setupRouter(db)
	clearDivisionData(db)

	requestBody := strings.NewReader(`{"division_name": "Testing Division"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5001/api/admindivisions", requestBody)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 409, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 409, int(responseBody["code"].(float64)))
	assert.Equal(t, "Conflict Duplicate", responseBody["status"])
	assert.Equal(t, "duplicate division name", responseBody["data"])
}

func TestUpdateAdminDivisionSuccess(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	divisionRepo := repository.NewAdminDivisionRepository(db)
	division := divisionRepo.Save(context.Background(), domain.AdminDivision{
		DivisionName: "Test1",
	})

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"division_name": "test1Update"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5001/api/admindivisions/"+strconv.Itoa(division.DivisionId), requestBody)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, division.DivisionId, int(responseBody["data"].(map[string]interface{})["division_id"].(float64)))
	assert.Equal(t, "test1Update", responseBody["data"].(map[string]interface{})["division_name"])
}

func TestUpdateAdminDivisionFailed(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	divisionRepo := repository.NewAdminDivisionRepository(db)
	division := divisionRepo.Save(context.Background(), domain.AdminDivision{
		DivisionName: "Test2",
	})

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"division_name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5001/api/admindivisions/"+strconv.Itoa(division.DivisionId), requestBody)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestDeleteAdminDivisionSuccess(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	divisionRepo := repository.NewAdminDivisionRepository(db)
	division := divisionRepo.Save(context.Background(), domain.AdminDivision{
		DivisionName: "Test3",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5001/api/admindivisions/"+strconv.Itoa(division.DivisionId), nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteAdminDivisionFailed(t *testing.T) {
	db := setupTestDB()

	router := setupRouter(db)
	clearDivisionData(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5001/api/admindivisions/404", nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestGetAdminDivisionSuccess(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	divisionRepo := repository.NewAdminDivisionRepository(db)
	division := divisionRepo.Save(context.Background(), domain.AdminDivision{
		DivisionName: "Test4",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admindivisions/"+strconv.Itoa(division.DivisionId), nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, division.DivisionId, int(responseBody["data"].(map[string]interface{})["division_id"].(float64)))
	assert.Equal(t, "Test4", responseBody["data"].(map[string]interface{})["division_name"])
}

func TestGetAdminDivisionFailed(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:5001/api/admindivisions/401", nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "Not Found", responseBody["status"])
}

func TestGetAllAdminDivisionSuccess(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	divisionRepo := repository.NewAdminDivisionRepository(db)
	division := divisionRepo.Save(context.Background(), domain.AdminDivision{
		DivisionName: "Test5",
	})

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admindivisions", nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	var divisions = responseBody["data"].([]interface{})
	divisionResponse := divisions[3].(map[string]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, division.DivisionId, int(divisionResponse["division_id"].(float64)))
	assert.Equal(t, "Test5", divisionResponse["division_name"])
}

func TestUnauthorizedApiKey(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admindivisions", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer xaas")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Invalid API Key", responseBody["status"])
}

func TestUnauthorizedBearer(t *testing.T) {
	db := setupTestDB()
	clearDivisionData(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admindivisions", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "Invalid Authorization Header", responseBody["status"])
}
