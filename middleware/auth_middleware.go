package middleware

import (
	"net/http"
	"strings"

	"github.com/emobodigo/golang_dashboard_api/helper"
	"github.com/emobodigo/golang_dashboard_api/model/payload"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-API-KEY") == "RAHASIA" {
		exludeAuthPath := []string{"/login", "/public"}
		shouldCheckBearer := true
		for _, path := range exludeAuthPath {
			if strings.Contains(r.URL.Path, path) {
				shouldCheckBearer = false
				break
			}
		}
		if shouldCheckBearer {
			authHeader := r.Header.Get("Authorization")
			// token := strings.TrimPrefix(authHeader, "Bearer ")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnauthorized)

				apiResponse := payload.ApiResponse{
					Code:   http.StatusUnauthorized,
					Status: "Invalid Authorization Header",
				}
				helper.WriteToResponseBody(w, apiResponse)
			} else {
				am.Handler.ServeHTTP(w, r)
			}
		} else {
			am.Handler.ServeHTTP(w, r)
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		apiResponse := payload.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "Invalid API Key",
		}
		helper.WriteToResponseBody(w, apiResponse)
	}

}
