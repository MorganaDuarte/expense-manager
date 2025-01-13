package controller

import (
	"encoding/json"
	"expense-manager/resource"
	"log"
	"net/http"
)

func GetBankAccountsByUserID(w http.ResponseWriter, r *http.Request) {
	database := resource.GetDatabaseInstance()
	defer database.Close()

	results, err := database.SelectBanksAccountsByUserID(1)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
