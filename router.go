package main

import (
	"expense-manager/controller"
	"net/http"
)

func registerRoutes() {
	http.HandleFunc("/api/save-bank-account", controller.SaveBankAccount)
	http.HandleFunc("/api/get-bank-accounts", controller.GetBankAccountsByUserID)
	http.HandleFunc("/api/values-received", controller.ValuesReceived)
}
