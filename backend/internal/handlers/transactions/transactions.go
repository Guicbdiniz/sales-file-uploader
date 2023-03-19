package transactions

import (
	"encoding/json"
	"guicbdiniz/hubla/backend/internal/models"
	"guicbdiniz/hubla/backend/internal/utils"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateTransactionsHandler returns an http handler with all the
// handlers of the /transactions route.
func CreateTransactionsHandler(models *models.Models) http.Handler {
	router := chi.NewRouter()

	addCreateTransactionsHandler(router, models)
	addGetAllTransactionHandler(router, models)

	return router
}

// addCreateTransactionsHandler adds a POST route to the passed mux
// to create new Transactions using the models.
//
// The request body must be an array of multiple Transactions.
func addCreateTransactionsHandler(router *chi.Mux, m *models.Models) {
	router.Post("/", func(res http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		var newTransactions []models.Transaction
		err = json.Unmarshal(body, &newTransactions)
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusBadRequest, err)
			return
		}

		for _, transaction := range newTransactions {
			err = m.Transactions.AddNewTransaction(transaction)
			if err != nil {
				utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
				return
			}

		}

		err = ProcessBalancesFromTransactions(m, newTransactions)
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		jsonResponse, err := utils.MarshalJsonResponse[[]models.Transaction](newTransactions)

		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		utils.SendJsonResponse(res, http.StatusCreated, jsonResponse)
	})
}

// addGetAllTransactionHandler adds a GET route to the passed mux
// to get all Transactions using the models.
func addGetAllTransactionHandler(router *chi.Mux, m *models.Models) {
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		transactions, err := m.Transactions.GetAllTransactions()
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		jsonResponse, err := utils.MarshalJsonResponse[[]models.Transaction](transactions)
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		utils.SendJsonResponse(res, http.StatusOK, jsonResponse)
	})
}
