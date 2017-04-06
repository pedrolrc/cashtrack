package cashcore

import (
     
)

type Expense struct {

}

func CreateExpenseAPI(description string) string {
    return "New expense created: " + description
}
