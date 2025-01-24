package controller

import (
	"encoding/json"
	"expense-manager/applicationservice"
	"expense-manager/resource"
	"log"
	"net/http"
)

func GetBankAccountsByUserID(w http.ResponseWriter, r *http.Request) {
	database := resource.GetDatabaseInstance()
	defer database.Close()

	results, err := applicationservice.GetBankAccountsByUser(1, database)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
