package resource

import (
	"context"
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

func (r *DatabaseResource) SaveValueReceived(value float32, date time.Time, description string, bank string) {
	sqlString := "INSERT INTO values_received (user_id, value, date, description, bank) VALUES($1, $2, $3, $4, $5)"

	_, err := r.Conn.Exec(context.Background(), sqlString, 1, value, date, description, bank)
	if err != nil {
		fmt.Println("Error saving value received:", err)
		return
	}
}
