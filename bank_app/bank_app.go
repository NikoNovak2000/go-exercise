package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt" // where the data is stored

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------------")
		// panic(err)
	}

	fmt.Println("Welcome to the Bank App!")
	fmt.Println("Reach us 24/7: ", randomdata.PhoneNumber())

	for {
		PresentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			}
		default:
			fmt.Println("Exited. Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			return
		}
	}
}
