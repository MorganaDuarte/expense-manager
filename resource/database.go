package resource

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DatabaseResource struct{}

var db *sql.DB
var databaseResource *DatabaseResource

func GetDatabaseInstance() *DatabaseResource {
	if db != nil {
		return databaseResource
	}

	cfg := mysql.Config{
		User:   "root",
		Passwd: "123456",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "expense_manager",
		Params: map[string]string{
			"allowNativePasswords": "true",
		},
	}
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	return databaseResource
}

func (r *DatabaseResource) SaveValueReceived(value float32, date time.Time, description string, bank string) {
	sqlString := "INSERT INTO values_received (user_id, value, date, description, bank) VALUES(?, ?, ?, ?, ?)"

	_, err := db.Exec(sqlString, 1, value, date, description, bank)
	if err != nil {
		fmt.Println("Error saving value received!")
		return
	}
}
