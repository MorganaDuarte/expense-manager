package main

import (
	"expense-manager/controller"
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/api/save-bank-account", controller.SaveBankAccount)
	http.HandleFunc("/api/get-bank-accounts", controller.GetBankAccountsByUserID)
	http.HandleFunc("/api/values-received", controller.ValuesReceived)

	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
