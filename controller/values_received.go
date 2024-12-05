package controller

import (
	"encoding/json"
	"expense-manager/resource"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ValuesReceivedInput struct {
	ValueReceived       float32 `json:"value_received"`
	DateReceived        string  `json:"date_received"`
	DescriptionReceived string  `json:"description_received"`
	AccountReceived     string  `json:"account_received"`
}

type ValuesReceivedResponse struct {
	Value       float32 `json:"value"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Account     string  `json:"account"`
}

func ValuesReceived(w http.ResponseWriter, r *http.Request) {
	var valueInput ValuesReceivedInput

	fmt.Println(valueInput)
	if err := json.NewDecoder(r.Body).Decode(&valueInput); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dateReceived, err := time.Parse("2006-01-02", valueInput.DateReceived)
	if err != nil {
		fmt.Println("Erro ao converter a string em data:", err)
		return
	}

	database := resource.GetDatabaseInstance()
	database.SaveValueReceived(valueInput.ValueReceived, dateReceived, valueInput.DescriptionReceived, valueInput.AccountReceived)
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
