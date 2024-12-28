package controller

import (
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

func SaveBankAccount(w http.ResponseWriter, r *http.Request) {
	var input *InputRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if validateSaveBankAccount(w, input) {
		return
	}

	database := resource.GetDatabaseInstance()
	database.Close()
	_, err := database.SaveBankAccount(input.BankValue, input.AcronymValue, input.AcronymValue)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func validateSaveBankAccount(w http.ResponseWriter, input *InputRequest) bool {
	if len(input.AcronymValue) > 3 {
		log.Println("Invalid AcronymValue: must be up to 3 letters")
		sendJSONError(w, "A sigla deve ter no máximo 3 letras.", http.StatusBadRequest)
		return true
	}
	return false
}
