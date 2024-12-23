package controller

import (
	"context"
	"encoding/json"
	"expense-manager/resource"
	"log"
	"net/http"
)

func GetAccountsByUserID(w http.ResponseWriter, r *http.Request) {
	database := resource.GetDatabaseInstance()
	results, err := database.SelectAccountsByUserID(1)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer database.Conn.Close(context.Background())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
