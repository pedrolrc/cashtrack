package cashcore

import (
)

type Expense struct {
    ID          string `json:"id,omitempty"`
    Date        string `json:"date,omitempty"`
    Description string `json:"description:omitempty"`
    Amount      string `json:"amount:omitempty"`
}

func CreateExpenseAPI(description string) string {
    return "New expense created: " + description
}
