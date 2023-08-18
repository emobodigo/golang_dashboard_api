package tests

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/emobodigo/golang_dashboard_api/model/domain"
	"github.com/emobodigo/golang_dashboard_api/repository"
	"github.com/stretchr/testify/assert"
)

func truncateTierData(db *sql.DB) {
	_, err := db.Exec("TRUNCATE admin_tier")
	if err != nil {
		panic(err)
	}
}

func TestCreateAdminTierSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"admin_level": 2, "division_id": 2, "level_title":"Test", "fulltime":1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5001/api/admintiers", requestBody)
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
	assert.Equal(t, float64(2), responseBody["data"].(map[string]interface{})["admin_level"])
}

func TestCreateAdminTierFailed(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"division_id": 2, "level_title":"Test", "fulltime":1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:5001/api/admintiers", requestBody)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(&responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "Bad Request", responseBody["status"])
}

func TestUpdateAdminTierSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)

	tierRepo := repository.NewAdminTierRepository(db)
	tier := tierRepo.Save(context.Background(), domain.AdminTier{
		AdminLevel: 2,
		DivisionId: 2,
		LevelTitle: "Test",
		Fulltime:   1,
	})
	fmt.Println(tier)

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"admin_level": 2, "division_id": 2, "level_title":"TestUpdate", "fulltime":1}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5001/api/admintiers/"+strconv.Itoa(int(tier.AdminTierId)), requestBody)
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
	assert.Equal(t, int(tier.AdminTierId), int(responseBody["data"].(map[string]interface{})["admin_tier_id"].(float64)))
	assert.Equal(t, "TestUpdate", responseBody["data"].(map[string]interface{})["level_title"])
}

func TestUpdateAdminTierFailed(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)

	tierRepo := repository.NewAdminTierRepository(db)
	tier := tierRepo.Save(context.Background(), domain.AdminTier{
		AdminLevel: 2,
		DivisionId: 2,
		LevelTitle: "Test",
		Fulltime:   1,
	})
	fmt.Println(tier)

	router := setupRouter(db)

	requestBody := strings.NewReader(`{"admin_level": 2, "division_id": 2, "level_title":"TestUpdate", "fulltime":1}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:5001/api/admintiers/404", requestBody)
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

func TestGetAdminTierSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)

	tierRepo := repository.NewAdminTierRepository(db)
	tier := tierRepo.Save(context.Background(), domain.AdminTier{
		AdminLevel: 2,
		DivisionId: 2,
		LevelTitle: "Test",
		Fulltime:   1,
	})
	fmt.Println(tier)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admintiers/"+strconv.Itoa(int(tier.AdminTierId)), nil)
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
	assert.Equal(t, int(tier.AdminTierId), int(responseBody["data"].(map[string]interface{})["admin_tier_id"].(float64)))
	assert.Equal(t, "Test", responseBody["data"].(map[string]interface{})["level_title"])
}

func TestGetAdminTierFailed(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)

	tierRepo := repository.NewAdminTierRepository(db)
	tier := tierRepo.Save(context.Background(), domain.AdminTier{
		AdminLevel: 2,
		DivisionId: 2,
		LevelTitle: "Test",
		Fulltime:   1,
	})
	fmt.Println(tier)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admintiers/404", nil)
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

func TestGetAllPageAdminTierSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)

	tierRepo := repository.NewAdminTierRepository(db)
	tier := tierRepo.Save(context.Background(), domain.AdminTier{
		AdminLevel: 2,
		DivisionId: 2,
		LevelTitle: "Test",
		Fulltime:   1,
	})
	fmt.Println(tier)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admintiers?page=1", nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	var tiers = responseBody["data"].([]interface{})
	tierResponse := tiers[0].(map[string]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, int(tier.AdminTierId), int(tierResponse["admin_tier_id"].(float64)))
	assert.Equal(t, "Test", tierResponse["level_title"])
}

func TestGetAllPageAdminTierFailed(t *testing.T) {
	db := setupTestDB()
	truncateTierData(db)

	tierRepo := repository.NewAdminTierRepository(db)
	tier := tierRepo.Save(context.Background(), domain.AdminTier{
		AdminLevel: 2,
		DivisionId: 2,
		LevelTitle: "Test",
		Fulltime:   1,
	})
	fmt.Println(tier)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:5001/api/admintiers?page=2", nil)
	setupHeaderWithAuth(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	fmt.Println(&responseBody)

	var tiers = responseBody["data"].([]interface{})

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, len(tiers), 0)
}
