package controller

import (
	"encoding/json"
	"expense-manager/applicationservice/bankaccount"
	"expense-manager/resource"
	"log"
	"net/http"
)

func SaveBankAccount(w http.ResponseWriter, r *http.Request) {
	var input *bankaccount.CreateBankAccountInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	database := resource.GetDatabaseInstance()
	defer database.Close()

	err := bankaccount.CreateBankAccount(input, database)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
