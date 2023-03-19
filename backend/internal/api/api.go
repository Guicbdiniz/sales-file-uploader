package api

import (
	"guicbdiniz/hubla/backend/internal/handlers/balances"
	"guicbdiniz/hubla/backend/internal/handlers/transactions"
	"guicbdiniz/hubla/backend/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Api contains the server logic (routes, handlers and models)
type Api struct {
	// Api multiplexer from the Chi package.
	mux *chi.Mux

	// Data access shared by all handlers.
	models *models.Models
}

// ServeHTTP is the single method of the http.Handler interface that makes
// the Api interoperable with the standard library.
func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	api.mux.ServeHTTP(w, r)
}

// CreateAPI returns a newly initialized Api object that implements
// the Handler interface.
func CreateAPI(models *models.Models) (*Api, error) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong"))
	})

	router.Mount("/transactions", transactions.CreateTransactionsHandler(models))
	router.Mount("/balances", balances.CreateBalancesHandler(models))

	return &Api{
		mux:    router,
		models: models,
	}, nil
}
