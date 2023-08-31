package testing_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"your-app-path/databases"
	"your-app-path/models"
	"your-app-path/websocket"
	"your-app-path/mahasiswacontroller"
)

func TestIndex(t *testing.T) {
	router := gin.Default()

	// Mock database connection
	databases.ConnectDatabaseForTesting()

	// Set up a mock WebSocket update function
	websocket.SetMockWebSocketUpdateFunction(func(message string) {})

	t.Run("Get All Students", func(t *testing.T) {
		// Create and insert mock Mahasiswa records into the database
		mockMahasiswas := []models.Mahasiswa{
			{ID: 1, Name: "John"},
			{ID: 2, Name: "Jane"},
		}
		databases.DB.Create(&mockMahasiswas)

		// Create a GET request
		req, _ := http.NewRequest("GET", "/mahasiswas", nil)
		rec := httptest.NewRecorder()

		// Serve the request
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string][]models.Mahasiswa
		err := mahasiswacontroller.ParseJSON(rec.Body.Bytes(), &response)
		assert.Nil(t, err)
		assert.NotNil(t, response["mahasiswas"])
		assert.Equal(t, len(mockMahasiswas), len(response["mahasiswas"]))
	})
}
