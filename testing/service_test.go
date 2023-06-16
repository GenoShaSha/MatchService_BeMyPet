package service_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"machingMicroService/model"
	"machingMicroService/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func TestMakeMatch(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Create the request body
	match := model.Match{
		ShelterID:     1,
		AnimalID:      123,
		AdopterID:     456,
		Picture:       "https://example.com/picture.jpg",
		FirstName:     "John",
		LastName:      "Doe",
		DateOfBirth:   "2000-01-01",
		Gender:        "Male",
		Type:          "Dog",
		Breed:         "Labrador Retriever",
		Shelter:       "ABC Shelter",
		Address:       "123 Main St",
		PostalCode:    "12345",
		Bio:           "A friendly and playful dog.",
		Status:        "Pending",
		AdopterEmail:  "test@example.com",
	}

	payload, err := json.Marshal(match)
	if err != nil {
		t.Fatal("Failed to marshal JSON payload:", err)
	}
	request := httptest.NewRequest(http.MethodPost, "/make-match", bytes.NewBuffer(payload))

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the MakeMatch handler function
	table := "matches_test" // Provide the table name
	router.POST("/make-match", func(c *gin.Context) {
		service.MakeMatch(c,table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Assert the response body
	var responseBody model.Match
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatal("Failed to unmarshal response body:", err)
	}

	// Perform additional assertions as needed
	// For example, you can assert that the returned match fields match the expected values
	assert.Equal(t, match.ShelterID, responseBody.ShelterID)
	assert.Equal(t, match.AnimalID, responseBody.AnimalID)
}

func TestGetMatches(t *testing.T) {
	// Set up the router
	router := gin.Default()
	
	request := httptest.NewRequest(http.MethodGet, "/getmatches", nil)

	// Create a test response recorder
	recorder := httptest.NewRecorder()
	// Call the CreateAnimal handler function
	table := "matches_test" // Provide the table name
	router.GET("/getmatches", func(c *gin.Context) {
		service.GetMatches(c, table)
	})
	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetAdopterMatches(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Define a sample adopter ID for testing
	adopterID := int64(456)

	// Create a JSON request body with the adopter ID
	requestBody := struct {
		AdopterID int64 `json:"adopter_id"`
	}{
		AdopterID: adopterID,
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a request with the JSON body
	request := httptest.NewRequest(http.MethodPost, "/getadoptermatches", bytes.NewReader(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the GetAdopterMatches handler function
	table := "matches_test" // Provide the table name
	router.POST("/getadoptermatches", func(c *gin.Context) {
		service.GetAdopterMatches(c, table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func TestGetShelterMatches(t *testing.T) {
	// Set up the router
	router := gin.Default()

	// Define a sample shelter ID for testing
	shelterID := int64(1)

	// Create a JSON request body with the shelter ID
	requestBody := struct {
		ShelterID int64 `json:"shelter_id"`
	}{
		ShelterID: shelterID,
	}

	// Convert the request body to JSON
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	// Create a request with the JSON body
	request := httptest.NewRequest(http.MethodPost, "/getsheltermatches", bytes.NewReader(jsonBody))
	request.Header.Set("Content-Type", "application/json")

	// Create a test response recorder
	recorder := httptest.NewRecorder()

	// Call the GetShelterMatches handler function
	table := "matches_test" // Provide the table name
	router.POST("/getsheltermatches", func(c *gin.Context) {
		service.GetShelterMatches(c, table)
	})

	// Perform the request
	router.ServeHTTP(recorder, request)

	// Assert the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)
}
