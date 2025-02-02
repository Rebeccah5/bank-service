package routes

import (
	"bank-service/internal/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(handler *handlers.AccountHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/balance", handler.GetBalance)
	r.Post("/deposit", handler.Deposit)
	r.Post("/withdraw", handler.Withdraw)

	return r
}
