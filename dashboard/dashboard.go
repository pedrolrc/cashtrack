package dashboard

import (
    "fmt"
	"net/http"
//    "encoding/json"
//	"github.com/gorilla/mux"
	"html/template"
)

type DefaultPage struct {
	Name string
}

func HandleDashboardIndex(respw http.ResponseWriter, req *http.Request) {
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

