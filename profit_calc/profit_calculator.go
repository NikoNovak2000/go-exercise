package main

import (
	"errors"
	"fmt"
	"os"
)

// Task 3
// 1) Validate user input --> show error msg & exit if invalid input is provided (no negative numbers, not 0)
// 2) Store calculated results into file

const resultsFile = "result.txt"

func main() {
	revenue, err := GetUserInput("Enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := GetUserInput("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := GetUserInput("Enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := CalculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
	storeResults(ebt, profit, ratio)
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile(resultsFile, []byte(results), 0644)
}

func GetUserInput(outputText string) (float64, error) {
	var userInput float64
	fmt.Print(outputText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}
	return userInput, nil
}

func CalculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}
