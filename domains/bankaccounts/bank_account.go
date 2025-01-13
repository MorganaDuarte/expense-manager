package bankaccounts

type BankAccount struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Acronym     string `json:"acronym"`
	Description string `json:"description"`
}

func New(id, userID int64, acronym, description string) *BankAccount {
	return &BankAccount{id, userID, acronym, description}
}
