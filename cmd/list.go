package cmd

import (
	"expense-tracker/storage"
	"fmt"
)

func ListExpenses() {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	fmt.Println("ID  Date       Description  Amount")
	for _, expense := range expenses {
		fmt.Printf("%d   %s   %s   $%.2f\n", expense.ID, expense.Date, expense.Description, expense.Amount)
	}
}
