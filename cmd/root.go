package cmd

import (
	"flag"
	"fmt"
	"os"
)


func Execute() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("description", "", "Expense description")
	amount := addCmd.Float64("amount", 0.0, "Expense amount")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "Expense ID to delete")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	month := summaryCmd.Int("month", 0, "Month number for summary (optional)")

	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *desc == "" || *amount <= 0 {
			fmt.Println("Error: Description and positive amount are required.")
			fmt.Println("Usage: expense-tracker add --description <description> --amount <amount>")
			return
		}
		AddExpense(*desc, *amount)
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *id <= 0 {
			fmt.Println("Error: Valid expense ID is required.")
			fmt.Println("Usage: expense-tracker delete --id <expense_id>")
			return
		}
		DeleteExpense(*id)
	case "list":
		listCmd.Parse(os.Args[2:])
		ListExpenses()
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		if *month > 0 {
			SummaryByMonth(*month)
		} else {
			Summary()
		}
	case "help":
		printHelp()
	default:
		fmt.Println("Error: Unknown command.")
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Usage: expense-tracker <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  add       Add a new expense")
	fmt.Println("            --description <description>  Description of the expense")
	fmt.Println("            --amount <amount>            Amount of the expense")
	fmt.Println("  delete    Delete an expense by ID")
	fmt.Println("            --id <expense_id>            ID of the expense to delete")
	fmt.Println("  list      List all expenses")
	fmt.Println("  summary   View a summary of expenses")
	fmt.Println("            --month <month_number>       (Optional) Month for the summary")
	fmt.Println("  help      Show this help message")
}
