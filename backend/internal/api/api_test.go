package api

import (
	"bytes"
	"encoding/json"
	"guicbdiniz/hubla/backend/internal/models"
	"guicbdiniz/hubla/backend/internal/utils"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	api, err := CreateAPI(nil)
	utils.CheckTestError(t, err, "error while creating api")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	api.ServeHTTP(w, req)

	utils.AssertEqual(t, http.StatusOK, w.Code, "Ping route did not return http.StatusOK")
}

func TestMultipleRequests(t *testing.T) {
	m, err := models.CreateTestModels()
	utils.CheckTestError(t, err, "error while creating test models")

	api, err := CreateAPI(m)
	utils.CheckTestError(t, err, "error while creating api")

	initialTransactions := []models.Transaction{
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Guilherme",
			Type:    "ProducerSale",
			Value:   100.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Luisa",
			Type:    "AffiliatedSale",
			Value:   100.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Luisa",
			Type:    "PaidCommission",
			Value:   10.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Guilherme",
			Type:    "ReceivedCommission",
			Value:   10.0,
		},
	}
	jsonBytes, err := json.Marshal(initialTransactions)
	utils.CheckTestError(t, err, "error captured creating json body")

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(jsonBytes))
	api.ServeHTTP(w, req)

	transactions, err := m.Transactions.GetAllTransactions()
	utils.CheckTestError(t, err, "error while getting all transactions")
	utils.AssertEqual(t, 4, len(transactions), "wrong number of transactions were created")

	clients, err := m.Clients.GetAllClients()
	utils.CheckTestError(t, err, "error while getting all clients")
	utils.AssertEqual(t, 2, len(clients), "wrong number of clients were created")

	secondaryTransactions := []models.Transaction{
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Guilherme",
			Type:    "ProducerSale",
			Value:   100.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Bruno",
			Type:    "AffiliatedSale",
			Value:   100.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Bruno",
			Type:    "PaidCommission",
			Value:   10.0,
		},
		models.Transaction{
			Date:    time.Now(),
			Product: "GoTutorial",
			Seller:  "Guilherme",
			Type:    "ReceivedCommission",
			Value:   10.0,
		},
	}

	jsonBytes, err = json.Marshal(secondaryTransactions)
	utils.CheckTestError(t, err, "error captured creating json body")

	req, _ = http.NewRequest(http.MethodPost, "/transactions", bytes.NewReader(jsonBytes))
	api.ServeHTTP(w, req)

	transactions, err = m.Transactions.GetAllTransactions()
	utils.CheckTestError(t, err, "error while getting all transactions")
	utils.AssertEqual(t, 8, len(transactions), "wrong number of transactions were created")

	clients, err = m.Clients.GetAllClients()
	utils.CheckTestError(t, err, "error while getting all clients")
	utils.AssertEqual(t, 3, len(clients), "wrong number of clients were created")
}
