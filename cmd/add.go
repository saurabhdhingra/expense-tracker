package cmd

import (
	"expense-tracker/models"
	"expense-tracker/storage"
	"fmt"
	"time"
)

func AddExpense(description string, amount float64) {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	newID := len(expenses) + 1
	currentDate := time.Now().Format("2006-01-02")
	newExpense := models.Expense{
		ID:          newID,
		Date:        currentDate,
		Description: description,
		Amount:      amount,
	}

	expenses = append(expenses, newExpense)
	err = storage.SaveExpenses(expenses)
	if err != nil {
		fmt.Println("Error saving expense:", err)
		return
	}

	fmt.Printf("Expense added successfully (ID: %d)\n", newID)
}
