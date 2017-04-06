package main

import (
	"fmt"
    "log"
	"net/http"
	"github.com/pedrolrc/cashtrack/cashcore"
	"html/template"
	"github.com/gorilla/mux"
    "encoding/json"
)

type DefaultPage struct {
	Name string
}

// Handling the URLs/Endpoints

func handleDashboardIndex(respw http.ResponseWriter, req *http.Request) {
	// URL / was requested
	fmt.Println("DEBUG :: Dashboard's Index called")

	username := "Um lesk muito louco"

	if req.Method == "POST" {
		fmt.Println("DEBUG :: Change Name submitted")
		req.ParseForm()

		username = req.Form["username"][0]
	}

	pageForIndex := &DefaultPage{ Name: username }

	tmp, _ := template.ParseFiles("templates/default.html")

	_ = tmp.Execute(respw, pageForIndex)
}

func handleCashCoreExpenseGet(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense Get API requested.")
    params := mux.Vars(req)

    fmt.Println("Getting " + params["id"])
    json.NewEncoder(respw).Encode(&cashcore.Expense{})
}

func handleCashCoreExpenseCreate(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense Create API requested.")
    params := mux.Vars(req)

    fmt.Println("Creating " + params["id"])
    json.NewEncoder(respw).Encode(&cashcore.Expense{})
}

func handleCashCoreExpenseDelete(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense Delete API requested.")
    params := mux.Vars(req)

    fmt.Println("Deleting " + params["id"])
    json.NewEncoder(respw).Encode(&cashcore.Expense{})
}

func main () {
    muxRouter := mux.NewRouter()

    // DASHBOARD - Home
    muxRouter.HandleFunc("/", handleDashboardIndex).Methods("GET")

    // CASHCORE - Expense
    muxRouter.HandleFunc("/expense/{id}", handleCashCoreExpenseGet).Methods("GET")
    muxRouter.HandleFunc("/expense/{id}", handleCashCoreExpenseCreate).Methods("POST")
    muxRouter.HandleFunc("/expense/{id}", handleCashCoreExpenseDelete).Methods("DELETE")

	fmt.Println("STATUS :: Cashtrack Webserver is up and listening for requests...")

    log.Fatal(http.ListenAndServe(":8080", muxRouter))
}
