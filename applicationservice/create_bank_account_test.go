package applicationservice_test

import (
	"expense-manager/applicationservice"
	"expense-manager/resource"
	"testing"
)

func TestCreateBankAccount(t *testing.T) {
	createBankAccountInput := &applicationservice.CreateBankAccountInput{
		Acronym:     "NuC",
		Description: "Conta Corrente",
	}

	inMemory := resource.GetInstance()
	defer inMemory.CleanMemory()

	err := applicationservice.CreateBankAccount(createBankAccountInput, inMemory)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

}

func TestCreateBankAccountWithAcronymEmpty(t *testing.T) {
	createBankAccountInput := &applicationservice.CreateBankAccountInput{
		Acronym:     "",
		Description: "Cartão",
	}

	inMemory := resource.GetInstance()
	defer inMemory.CleanMemory()

	err := applicationservice.CreateBankAccount(createBankAccountInput, inMemory)
	expectedMessage := "A sigla não pode ser vazia"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
}

func TestCreateBankAccountWithAcronymTooLong(t *testing.T) {
	createBankAccountInput := &applicationservice.CreateBankAccountInput{
		Acronym:     "aeeeeeee",
		Description: "Cartão",
	}

	inMemory := resource.GetInstance()
	defer inMemory.CleanMemory()

	err := applicationservice.CreateBankAccount(createBankAccountInput, inMemory)
	expectedMessage := "A sigla deve ter no máximo 3 letras"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
}

func TestCreateBankAccountWithRepeatedAcronym(t *testing.T) {
	createBankAccountInput := &applicationservice.CreateBankAccountInput{
		Acronym:     "ae",
		Description: "Cartão",
	}
	createBankAccountInput2 := &applicationservice.CreateBankAccountInput{
		Acronym:     "ae",
		Description: "Cartão",
	}
	inMemory := resource.GetInstance()
	defer inMemory.CleanMemory()

	applicationservice.CreateBankAccount(createBankAccountInput, inMemory)
	err := applicationservice.CreateBankAccount(createBankAccountInput2, inMemory)
	expectedMessage := "sigla já existente"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
}
