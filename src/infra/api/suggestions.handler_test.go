package api_test

import (
	"ifood-backend-test/src/infra/api"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestApi(t *testing.T) {
	err := godotenv.Load("../../../.env"); if err != nil {
		log.Println("Error on load env file")
	}

	r := chi.NewRouter()
  api.MakeSuggestionHandler(r)

	req, err := http.NewRequest("GET", "/api/latitude/-19.7762/longitude/-43.8819", nil)
  require.Nil(t, err)

	// Record the response
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Get the status code
	statusCode := rr.Code

	// Get the response body
	// responseBody := rr.Body.String()

	// Optionally, you can include assertions
	require.Equal(t, http.StatusOK, statusCode)
}