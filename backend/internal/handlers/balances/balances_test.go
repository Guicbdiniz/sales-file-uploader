package balances

import (
	"guicbdiniz/hubla/backend/internal/models"
	"guicbdiniz/hubla/backend/internal/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBalancesHandler(t *testing.T) {
	m, err := models.CreateTestModels()
	utils.CheckTestError(t, err, "error while creating test models")

	handler := CreateBalancesHandler(m)

	testClient := models.Client{
		IsProducer: false,
		Name:       "Guilherme",
	}
	m.Clients.AddNewClient(testClient)
	testGetAllBalances(t, handler)
}

func testGetAllBalances(t *testing.T, handler http.Handler) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusOK, response.StatusCode, "GET request to /balances did not return the correct status")

	body, err := io.ReadAll(response.Body)
	utils.CheckTestError(t, err, "error captured while reading a response")

	jsonBody, err := utils.UnmarshalJsonResponse[[]models.Client](body)
	utils.CheckTestError(t, err, "error captured while unsmarshiling balances")

	utils.AssertEqual(t, 1, len(jsonBody.Data), "GET request to /balances did not return correct body")
	utils.AssertEqual(t, jsonBody.ErrorText, "", "GET request to /balances did not return correct error text")
	utils.AssertEqual(t, jsonBody.Data[0].Name, "Guilherme", "GET request to /balances did not return correct error text")
}
