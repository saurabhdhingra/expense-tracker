package cmd

import (
	"expense-tracker/storage"
	"fmt"
)

func ExportToCSV(filename string) {
	err := storage.ExportExpensesToCSV(filename)
	if err != nil {
		fmt.Println("Error exporting expenses to CSV:", err)
		return
	}

	fmt.Printf("Expenses exported successfully to %s\n", filename)
}
