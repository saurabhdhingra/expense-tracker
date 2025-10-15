# Expense Tracker CLI

A simple command-line application to track expenses efficiently. Users can add, update, delete, and view their expenses, along with summaries of their spending habits.

## Features
- Add an expense with a description and amount.
- Update an existing expense.
- Delete an expense.
- View all expenses.
- View a summary of total expenses.
- View expenses for a specific month.
- Export expenses to CSV.

## Installation
```sh
git clone https://github.com/saurabhdhingra/expense-tracker.git
cd expense-tracker
go build -o expense-tracker
```

## Usage

### Add an Expense
```sh
./expense-tracker add --description "Lunch" --amount 20
```

### List Expenses
```sh
./expense-tracker list
```

### View Summary
```sh
./expense-tracker summary
```

### Delete an Expense
```sh
./expense-tracker delete --id 1
```

### Summary by Month
```sh
./expense-tracker summary --month 8
```

## Error Handling
- Ensures only positive amounts are added.
- Prevents deletion of non-existent expenses.
- Validates user input for month filtering.

## Inspired by roadmap.sh
- https://roadmap.sh/projects/expense-tracker