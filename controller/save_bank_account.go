package controller

import (
	"encoding/json"
	"expense-manager/resource"
	"fmt"
	"log"
	"net/http"
)

type SaveBankAccountInput struct {
	AcronymValue     string `json:"acronymValue"`
	DescriptionValue string `json:"descriptionValue"`
}

func SaveBankAccount(w http.ResponseWriter, r *http.Request) {
	var input *SaveBankAccountInput

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println("Error decoding JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := validateSaveBankAccountInput(input)
	if err != nil {
		sendJSONError(w, err, http.StatusBadRequest)
		return
	}

	database := resource.GetDatabaseInstance()
	defer database.Close()
	err = database.SaveBankAccount(input.AcronymValue, input.DescriptionValue)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func validateSaveBankAccountInput(input *SaveBankAccountInput) error {
	if len(input.AcronymValue) == 0 {
		log.Println("AcronymValue can`t be empty")
		return fmt.Errorf("A sigla não pode ser vazia")
	}

	if len(input.AcronymValue) > 3 {
		log.Println("Invalid AcronymValue: must be up to 3 letters")
		return fmt.Errorf("A sigla deve ter no máximo 3 letras")
	}
	return nil
}
