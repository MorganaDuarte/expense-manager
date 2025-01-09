package bankaccount

import (
	"expense-manager/domains/bankaccounts"
	"expense-manager/resource"
	"log"
)

func GetBankAccountsByUser(id int, resource resource.Interface) ([]*bankaccounts.BankAccount, error) {
	results, err := resource.SelectBanksAccountsByUserID(int64(id))
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	return results, nil
}
