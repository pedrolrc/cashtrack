package cashcore

import (
    "fmt"
	"net/http"
    "encoding/json"
	"github.com/gorilla/mux"
)

type Expense struct {
    ID          string `json:"id,omitempty"`
    Date        string `json:"date,omitempty"`
    Description string `json:"description,omitempty"`
    Amount      string `json:"amount,omitempty"`
}

func HandleCashCoreExpenseGet(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense Get API requested.")
    params := mux.Vars(req)

    fmt.Println("Getting " + params["id"])
    json.NewEncoder(respw).Encode(&Expense{})
}

func HandleCashCoreExpenseCreate(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense Create API requested.")
    params := mux.Vars(req)

    fmt.Println("Creating " + params["id"])
    json.NewEncoder(respw).Encode(&Expense{})
}

func HandleCashCoreExpenseDelete(respw http.ResponseWriter, req *http.Request) {
    fmt.Println("Cashcore's Expense Delete API requested.")
    params := mux.Vars(req)

    fmt.Println("Deleting " + params["id"])
    json.NewEncoder(respw).Encode(&Expense{})
}

