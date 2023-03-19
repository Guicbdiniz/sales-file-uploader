package balances

import (
	"guicbdiniz/hubla/backend/internal/models"
	"guicbdiniz/hubla/backend/internal/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateBalancesHandler returns an http handler with all the
// handlers of the /balances route.
//
// Balances are basically the same as clients.
func CreateBalancesHandler(models *models.Models) http.Handler {
	router := chi.NewRouter()

	addGetAllBalancesHandler(router, models)

	return router
}

// addGetAllBalancesHandler adds a GET route to the passed mux
// to get all Balances using the models.
func addGetAllBalancesHandler(router *chi.Mux, m *models.Models) {
	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		clients, err := m.Clients.GetAllClients()
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		jsonResponse, err := utils.MarshalJsonResponse[[]models.Client](clients)
		if err != nil {
			utils.SendJsonErrorResponse(res, http.StatusInternalServerError, err)
			return
		}

		utils.SendJsonResponse(res, http.StatusOK, jsonResponse)
	})
}
