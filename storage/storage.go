package storage

import (
	"encoding/json"
	"errors"
	"expense-tracker/models"
	"os"
)

var dataFile = "expenses.json"

func LoadExpenses() ([]models.Expense, error) {
	var expenses []models.Expense
	file, err := os.ReadFile(dataFile)
	if err != nil {
		return expenses, err
	}
	json.Unmarshal(file, &expenses)
	return expenses, nil
}

func SaveExpenses(expenses []models.Expense) error {
	data, err := json.MarshalIndent(expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func DeleteExpense(id int) error {
	expenses, err := LoadExpenses()
	if err != nil {
		return err
	}

	found := false
	for i, exp := range expenses {
		if exp.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("expense ID not found")
	}

	return SaveExpenses(expenses)
}
