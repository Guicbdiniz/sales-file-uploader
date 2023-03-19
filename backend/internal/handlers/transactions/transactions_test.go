package transactions

import (
	"bytes"
	"encoding/json"
	"guicbdiniz/hubla/backend/internal/models"
	"guicbdiniz/hubla/backend/internal/utils"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestBalancesFromTransactions(t *testing.T) {
	m, err := models.CreateTestModels()
	utils.CheckTestError(t, err, "error while creating test models")

	newTransactions := []models.Transaction{
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "John",
			Type:    "ProducerSale",
			Value:   50.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Maria",
			Type:    "AffiliatedSale",
			Value:   50.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Maria",
			Type:    "PaidCommission",
			Value:   10.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "John",
			Type:    "ReceivedCommission",
			Value:   10.0,
		},
	}
	err = ProcessBalancesFromTransactions(m, newTransactions)
	utils.CheckTestError(t, err, "error while processing balances from transactions")

	clients, err := m.Clients.GetAllClients()
	utils.CheckTestError(t, err, "error while getting all clients")

	utils.AssertEqual(t, 2, len(clients), "wrong number of clients were created")
	utils.AssertEqual(t, 40.0, clients[1].Balance, "wrong balance inserted into Maria")
	utils.AssertEqual(t, 60.0, clients[0].Balance, "wrong balance inserted into John")
}

func TestTransactionsHandler(t *testing.T) {
	models, err := models.CreateTestModels()
	utils.CheckTestError(t, err, "error while creating test models")

	handler := CreateTransactionsHandler(models)

	testAddTransactions(t, handler)
	testGetAllTransactions(t, handler)
	testClientWasCreated(t, models)
}

func testAddTransactions(t *testing.T, handler http.Handler) {
	transaction := models.Transaction{
		Date:    time.Now(),
		Product: "GoTutorial",
		Seller:  "Myself",
		Type:    "ProducerSale",
	}
	jsonBytes, err := json.Marshal([]models.Transaction{transaction})
	utils.CheckTestError(t, err, "error captured creating json body")

	request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(jsonBytes))
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusCreated, response.StatusCode, "POST request to /transactions did not return the correct status")
}

func testGetAllTransactions(t *testing.T, handler http.Handler) {
	request := httptest.NewRequest(http.MethodGet, "/", nil)
	responseRecorder := httptest.NewRecorder()

	handler.ServeHTTP(responseRecorder, request)
	response := responseRecorder.Result()

	utils.AssertEqual(t, http.StatusOK, response.StatusCode, "GET request to /transactions did not return the correct status")

	body, err := io.ReadAll(response.Body)
	utils.CheckTestError(t, err, "error captured while reading a response")

	jsonBody, err := utils.UnmarshalJsonResponse[[]models.Transaction](body)
	utils.CheckTestError(t, err, "error captured while unsmarshiling transactions")

	utils.AssertEqual(t, 1, len(jsonBody.Data), "GET request to /transactions did not return correct body")
	utils.AssertEqual(t, jsonBody.ErrorText, "", "GET request to /transactions did not return correct error text")
	utils.AssertEqual(t, "GoTutorial", jsonBody.Data[0].Product, "GET request to /transactions did not return correct body")
}

func testClientWasCreated(t *testing.T, models *models.Models) {
	clients, err := models.Clients.GetAllClients()
	utils.CheckTestError(t, err, "error while getting all clients")

	utils.AssertEqual(t, len(clients), 1, "POST request to /transactions did not created client as expected")
}
