package applicationservice

import (
	"expense-manager/domains/bankaccounts"
	"expense-manager/resource"
)

func GetBankAccountsByUser(id int, resource resource.Interface) ([]*bankaccounts.BankAccount, error) {
	results, err := resource.SelectBanksAccountsByUserID(id)
	if err != nil {
		return nil, err
	}

	return results, nil
}
