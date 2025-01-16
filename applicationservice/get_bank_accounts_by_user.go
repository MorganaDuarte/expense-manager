package applicationservice

import (
	"expense-manager/domains/bankaccount"
	"expense-manager/resource"
)

func GetBankAccountsByUser(id int, resource resource.Interface) ([]*bankaccount.BankAccount, error) {
	results, err := resource.SelectBanksAccountsByUserID(id)
	if err != nil {
		return nil, err
	}

	return results, nil
}
