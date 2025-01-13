package bankaccount_test

import (
	"expense-manager/applicationservice/bankaccount"
	"expense-manager/resource"
	"testing"
)

func TestCreateBankAccount(t *testing.T) {
	createBankAccountInput := &bankaccount.CreateBankAccountInput{
		Acronym:     "NuC",
		Description: "Conta Corrente",
	}

	inMemory := resource.GetInstance()
	err := bankaccount.CreateBankAccount(createBankAccountInput, inMemory)
	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}

	inMemory.CleanMemory()
}

func TestCreateBankAccountWithAcronymEmpty(t *testing.T) {
	createBankAccountInput := &bankaccount.CreateBankAccountInput{
		Acronym:     "",
		Description: "Cartão",
	}

	inMemory := resource.GetInstance()
	err := bankaccount.CreateBankAccount(createBankAccountInput, inMemory)

	expectedMessage := "A sigla não pode ser vazia"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
	inMemory.CleanMemory()
}

func TestCreateBankAccountWithAcronymTooLong(t *testing.T) {
	createBankAccountInput := &bankaccount.CreateBankAccountInput{
		Acronym:     "aeeeeeee",
		Description: "Cartão",
	}

	inMemory := resource.GetInstance()
	err := bankaccount.CreateBankAccount(createBankAccountInput, inMemory)

	expectedMessage := "A sigla deve ter no máximo 3 letras"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
	inMemory.CleanMemory()
}
