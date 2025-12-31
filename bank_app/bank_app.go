package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func GetBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000, errors.New("Failed to find balance file.")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value.")
	}

	return balance, nil
}

func WriteBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644) // 0644, way of encoding file permissions --> this set means the owner of the file will be able to read and write the file, while others just read
}

func CheckAccountBalance(balance float64) {
	fmt.Println("Your balance is: ", balance)
}

func BankOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func main() {
	var accountBalance, err = GetBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic(err)
	}

	fmt.Println("Welcome to the Bank App!")
	for {
		BankOptions()
		var userChoice int
		fmt.Print("Please input the number of the choosen action(1-4): ")
		fmt.Scan(&userChoice)
		fmt.Println("Your choice: ", userChoice)

		switch userChoice {
		case 1:
			CheckAccountBalance(accountBalance)
		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			CheckAccountBalance(accountBalance)
			fmt.Print("Writing to file!")
			WriteBalanceToFile(accountBalance)
		case 3:
			CheckAccountBalance(accountBalance)
			fmt.Print("Amount of money you wish to withdraw: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			var checkAmount float64 = accountBalance - withdrawalAmount

			if withdrawalAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}

			if checkAmount < 0 {
				fmt.Println("Cannot withdraw more money than you have.")
				continue
			} else {
				fmt.Println("You successfully withdraw this amount: ", withdrawalAmount)
				accountBalance -= withdrawalAmount
				CheckAccountBalance(accountBalance)
				WriteBalanceToFile(accountBalance)
			}
		default:
			fmt.Println("Exited. Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
