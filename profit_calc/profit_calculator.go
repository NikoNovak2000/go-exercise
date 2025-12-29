package main

import "fmt"

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Enter Revenue: ")  Replace this code with a custom function that outputs the text, waits for user to enter some text, return the entered value
	// fmt.Scan(&revenue)            In main, call this new made function and store it in one of the variables

	revenue := getUserInput("Enter Revenue: ")

	expenses := getUserInput("Enter Expenses: ")

	taxRate := getUserInput("Enter Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)

}

func getUserInput(outputText string) float64 {
	var userInput float64
	fmt.Print(outputText)
	fmt.Scan(&userInput)
	return userInput
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
