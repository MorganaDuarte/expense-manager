package controller

import (
	"context"
	"encoding/json"
	"expense-manager/resource"
	"log"
	"net/http"
)

type InputRequest struct {
	BankValue    string `json:"bank_value"`
	AccountValue string `json:"account_value"`
	AcronymValue string `json:"acronym_value"`
}

func SaveAccount(w http.ResponseWriter, r *http.Request) {
	var input *InputRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database := resource.GetDatabaseInstance()
	defer database.Conn.Close(context.Background())
	_, err := database.SaveAccount(input.BankValue, input.AcronymValue, input.AcronymValue)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
