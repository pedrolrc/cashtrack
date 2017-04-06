package main

import (
	"fmt"
	"net/http"
	"github.com/pedrolrc/cashtrack/cashcore"
	"html/template"
)

type DefaultPage struct {
	Name string
}

// Handling the URLs/Endpoints

func handleIndex(respw http.ResponseWriter, req *http.Request) {
	// URL / was requested
	fmt.Println("DEBUG :: handleIndex called")

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

func handleCashcoreExpenseGeneric(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense APIs requested.")
    fmt.Println(cashcore.CreateExpenseAPI("chevron gas"))
}

func main () {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/expense", handleCashcoreExpenseGeneric)

	fmt.Println("STATUS :: Cashtrack Webserver is up and listening for requests...")

	http.ListenAndServe(":8080", nil)
}
