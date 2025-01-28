package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func ExportExpensesToCSV(filename string) error {
	expenses, err := LoadExpenses()
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()


	writer.Write([]string{"ID", "Date", "Description", "Amount"})


	for _, expense := range expenses {
		writer.Write([]string{
			strconv.Itoa(expense.ID),
			expense.Date,
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
		})
	}

	return nil
}
