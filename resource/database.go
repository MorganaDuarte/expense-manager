package resource

import (
	"context"
	"expense-manager/domains/bankaccounts"
	"fmt"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"time"
)

type DatabaseResource struct {
	Conn *pgx.Conn
}

var dbResource *DatabaseResource

func GetDatabaseInstance() *DatabaseResource {
	if dbResource != nil {
		return dbResource
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE"))
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	dbResource = &DatabaseResource{
		Conn: conn,
	}
	return dbResource
}

func (r *DatabaseResource) Close() {
	dbResource = nil
	r.Conn.Close(context.Background())
}

func (r *DatabaseResource) SaveValueReceived(value float32, date time.Time, description string, bank string) {
	sqlString := "INSERT INTO values_received (user_id, value, date, description, bank) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Conn.Exec(context.Background(), sqlString, 1, value, date, description, bank)
	if err != nil {
		fmt.Println("Error saving value received:", err)
		return
	}
}

func (r *DatabaseResource) SaveBankAccount(acronym, description string) (int64, error) {
	sqlString := "INSERT INTO bank_accounts (user_id, acronym, description) VALUES($1, $2, $3)"

	result, err := r.Conn.Exec(context.Background(), sqlString, 1, acronym, description)
	if err != nil {
		fmt.Println("Error saving account", err)
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	return rowsAffected, nil
}

func (r *DatabaseResource) SelectBanksAccountsByUserID(id int64) ([]*bankaccounts.BankAccount, error) {
	sqlString := "SELECT * FROM bank_accounts WHERE user_id = $1"

	response, err := r.Conn.Query(context.Background(), sqlString, id)
	if err != nil {
		return nil, err
	}

	var results []*bankaccounts.BankAccount
	for response.Next() {
		var id, userID int64
		var acronym, description string

		err = response.Scan(&id, &userID, &acronym, &description)
		if err != nil {
			return nil, err
		}

		row := bankaccounts.New(id, userID, acronym, description)
		results = append(results, row)
	}

	return results, nil
}
