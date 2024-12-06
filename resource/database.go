package resource

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "morganaborba"
	password = "123456"
	dbname   = "expense_manager"
	sslmode  = false
)

type DatabaseResource struct{}

var db *sql.DB
var databaseResource *DatabaseResource

func GetDatabaseInstance() *DatabaseResource {
	fmt.Printf("Accessing %s ... ", dbname)
	if db != nil {
		return databaseResource
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return databaseResource
}

func (r *DatabaseResource) SaveValueReceived(value float32, date time.Time, description string, bank string) {
	sqlString := "INSERT INTO values_received (user_id, value, date, description, bank) VALUES($1, $2, $3, $4, $5)"

	_, err := db.Exec(sqlString, 1, value, date, description, bank)
	if err != nil {
		fmt.Println("Error saving value received:", err)
		return
	}
}
