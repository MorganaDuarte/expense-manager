package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type ValuesReceivedInput struct {
	ValueReceived       int    `json:"value_received"`
	DateReceived        string `json:"date_received"`
	DescriptionReceived string `json:"description_received"`
	AccountReceived     string `json:"account_received"`
}

type ValuesReceivedResponse struct {
	Value       int    `json:"value"`
	Date        string `json:"date"`
	Description string `json:"description"`
	Account     string `json:"account"`
}

func ValuesReceived(w http.ResponseWriter, r *http.Request) {
	var valueInput ValuesReceivedInput

	if err := json.NewDecoder(r.Body).Decode(&valueInput); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := ValuesReceivedResponse{
		Value:       valueInput.ValueReceived,
		Date:        valueInput.DateReceived,
		Description: valueInput.DescriptionReceived,
		Account:     valueInput.AccountReceived,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
