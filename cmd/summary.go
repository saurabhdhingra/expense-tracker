package cmd

import (
	"expense-tracker/storage"
	"fmt"
	"strconv"
	"time"
)

func Summary() {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	total := 0.0
	for _, expense := range expenses {
		total += expense.Amount
	}
	fmt.Printf("Total expenses: $%.2f\n", total)
}

func SummaryByMonth(month int) {
	expenses, err := storage.LoadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}

	if month < 1 || month > 12 {
		fmt.Println("Error: Invalid month. Please provide a month between 1 and 12.")
		return
	}

	total := 0.0
	currentYear := strconv.Itoa(time.Now().Year())
	for _, expense := range expenses {
		if len(expense.Date) >= 7 && expense.Date[:4] == currentYear {
			expenseMonth, _ := strconv.Atoi(expense.Date[5:7])
			if expenseMonth == month {
				total += expense.Amount
			}
		}
	}
	fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(month).String(), total)
}