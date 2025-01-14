package bankaccounts

type BankAccount struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Acronym     string `json:"acronym"`
	Description string `json:"description"`
}

func New(id, userID int, acronym, description string) *BankAccount {
	return &BankAccount{id, userID, acronym, description}
}
