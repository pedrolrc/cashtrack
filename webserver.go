package main

import (
	"fmt"
	"net/http"
//	"github.com/pedrolrc/cashtrack/track"
	"html/template"
)

type DefaultPage struct {
	Name string
}

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

func handleLogCost(respw http.ResponseWriter, req *http.Request) {
	// URL /logcost was requested
	fmt.Println("DEBUG :: handleLogCost called")	
}

func main () {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/logcost", handleLogCost)
	
	fmt.Println("STATUS :: Cashtrack Webserver is up and listening for requests...")
	
	http.ListenAndServe(":8080", nil)
}
