package handler

import (
	"encoding/json"
	"net/http"
	"study-case/models"
	"study-case/service"
)

// JSON response utility
func JsonResponse(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// Create transaction handler
func CreateTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var t models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate transaction data
	if t.ContractNumber == "" || t.CustomerID <= 0 || t.OTR <= 0 {
		http.Error(w, "Invalid transaction data", http.StatusBadRequest)
		return
	}

	tenor := 12 // Example tenor, replace with dynamic value
	if err := service.CreateTransaction(t, tenor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	JsonResponse(w, map[string]string{"message": "Transaction created successfully"}, http.StatusCreated)
}
