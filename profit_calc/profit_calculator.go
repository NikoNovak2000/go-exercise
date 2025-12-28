package main

import "fmt"

// Build a Profit Calculator --> command line tool built with GO
// What does it do? --> asks users for revenue, expenses, and tax rate input
// Calculates earnings before tax(EBT) and earnings after tax(profit), as well as the ratio (EBT/profit)
// Output EBT, profit and the ratio
// EBT = Revenue - Expensees
// Profit = EBT * (1 - Tax Rate)
// ratio = ebt/profit

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter Tax Rate: ")
	fmt.Scan(&taxRate)

	EBT := revenue - expenses
	profit := EBT * (1 - taxRate/100)
	ratio := EBT / profit

	fmt.Printf("Earnings before tax (EBT): %v\n", EBT)
	fmt.Printf("Profit after tax: %v\n", profit)
	fmt.Printf("EBT to Profit ratio: %v\n", ratio)
}
