package handlers

import (
	"bank-service/internal/repository"
	"encoding/json"
	"net/http"
)

type AccountHandler struct {
	Repo *repository.AccountRepository
}

// NewAccountHandler initializes the handler
func NewAccountHandler(repo *repository.AccountRepository) *AccountHandler {
	return &AccountHandler{Repo: repo}
}

// GetBalance handles GET /balance requests
func (h *AccountHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	balance, err := h.Repo.GetBalance()
	if err != nil {
		http.Error(w, "failed to fetch balance", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"balance": balance})
}

// Deposit handles POST /deposit requests
func (h *AccountHandler) Deposit(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if request.Amount <= 0 {
		http.Error(w, "amount must be greater than zero", http.StatusBadRequest)
		return
	}

	err := h.Repo.Deposit(request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Deposit successful"})
}

// Withdraw handles POST /withdraw requests
func (h *AccountHandler) Withdraw(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Amount float64 `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if request.Amount <= 0 {
		http.Error(w, "amount must be greater than zero", http.StatusBadRequest)
		return
	}

	err := h.Repo.Withdraw(request.Amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "withdrawal successful"})
}
