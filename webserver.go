package main

import (
	"fmt"
    "log"
	"net/http"
	"github.com/pedrolrc/cashtrack/cashcore"
	"github.com/pedrolrc/cashtrack/dashboard"
	"github.com/gorilla/mux"
)

// Handling the URLs/Endpoints

func main () {
    muxRouter := mux.NewRouter()

    // DASHBOARD - Home
    muxRouter.HandleFunc("/", dashboard.HandleDashboardIndex).Methods("GET")

    // CASHCORE - Expense
    muxRouter.HandleFunc("/expense/{id}", cashcore.HandleCashCoreExpenseGet).Methods("GET")
    muxRouter.HandleFunc("/expense/{id}", cashcore.HandleCashCoreExpenseCreate).Methods("POST")
    muxRouter.HandleFunc("/expense/{id}", cashcore.HandleCashCoreExpenseDelete).Methods("DELETE")

	fmt.Println("STATUS :: Cashtrack Webserver is up and listening for requests...")

    log.Fatal(http.ListenAndServe(":8080", muxRouter))
}
