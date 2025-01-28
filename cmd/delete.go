package cmd

import (
	"expense-tracker/storage"
	"fmt"
)

func DeleteExpense(id int) {
	expenses, err := storage.LoadExpenses()


	found := false
	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Expense with ID %d not found.\n", id)
		return
	}

	err = storage.SaveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}

	fmt.Printf("Expense with ID %d deleted successfully.\n", id)
}
