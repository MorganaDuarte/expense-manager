package bankaccount_test

import (
	"expense-manager/applicationservice/bankaccount"
	"testing"
)

func TestCreateBankAccount(t *testing.T) {
	createBankAccountInput := &bankaccount.CreateBankAccountInput{
		AcronymValue:     "NuC",
		DescriptionValue: "Conta Corrente",
	}

	err := bankaccount.CreateBankAccount(createBankAccountInput)

	if err != nil {
		t.Errorf("expected no error, but got %v", err)
	}
}

func TestCreateBankAccountWithAcronymEmpty(t *testing.T) {
	createBankAccountInput := &bankaccount.CreateBankAccountInput{
		AcronymValue:     "",
		DescriptionValue: "Cartão",
	}

	err := bankaccount.CreateBankAccount(createBankAccountInput)

	expectedMessage := "A sigla não pode ser vazia"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
}

func TestCreateBankAccountWithAcronymTooLong(t *testing.T) {
	createBankAccountInput := &bankaccount.CreateBankAccountInput{
		AcronymValue:     "aeeeeeee",
		DescriptionValue: "Cartão",
	}

	err := bankaccount.CreateBankAccount(createBankAccountInput)

	expectedMessage := "A sigla deve ter no máximo 3 letras"
	if err == nil || err.Error() != expectedMessage {
		t.Errorf("expected error message %q, got %q", expectedMessage, err)
	}
}
