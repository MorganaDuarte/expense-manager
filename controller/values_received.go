package controller

import (
	"encoding/json"
	"log"
	"net/http"
)

type ValuesReceivedInput struct {
	ValueReceived       string `json:"value_received"`
	DateReceived        string `json:"date_received"`
	DescriptionReceived string `json:"description_received"`
	AccountReceived     string `json:"account_received"`
}

func ValuesReceived(w http.ResponseWriter, r *http.Request) {
	var valueResponse ValuesReceivedInput

	if err := json.NewDecoder(r.Body).Decode(&valueResponse); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"value":       valueResponse.ValueReceived,
		"date":        valueResponse.DateReceived,
		"description": valueResponse.DescriptionReceived,
		"account":     valueResponse.AccountReceived,
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
