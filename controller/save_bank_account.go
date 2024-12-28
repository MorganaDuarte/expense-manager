package controller

import (
	"encoding/json"
	"expense-manager/resource"
	"log"
	"net/http"
)

type InputRequest struct {
	AcronymValue     string `json:"acronym_value"`
	DescriptionValue string `json:"description_value"`
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
	defer database.Close()
	_, err := database.SaveBankAccount(input.AcronymValue, input.DescriptionValue)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}

func validateSaveBankAccount(w http.ResponseWriter, input *InputRequest) bool {
	if len(input.AcronymValue) == 0 {
		log.Println("AcronymValue can`t be empty")
		sendJSONError(w, "A sigla não pode ser vazia.", http.StatusBadRequest)
		return true
	}

	if len(input.AcronymValue) > 3 {
		log.Println("Invalid AcronymValue: must be up to 3 letters")
		sendJSONError(w, "A sigla deve ter no máximo 3 letras.", http.StatusBadRequest)
		return true
	}
	return false
}
