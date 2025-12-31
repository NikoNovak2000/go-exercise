package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to the Bank App!")

	for {
		bankOptions()
		var userChoice int
		fmt.Print("Please input the number of the choosen action(1-4): ")
		fmt.Scan(&userChoice)
		fmt.Println("Your choice: ", userChoice)

		// Can store checks in variables and get a value of type boolean
		// wantsCheckBalance := userChoice == 1

		if userChoice == 1 {
			checkAccountBalance(accountBalance)
		} else if userChoice == 2 {
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			checkAccountBalance(accountBalance)
		} else if userChoice == 3 {
			checkAccountBalance(accountBalance)
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
				checkAccountBalance(accountBalance)
			}
		} else if userChoice == 4 {
			fmt.Println("Exited. Goodbye!")
			break
		} else {
			fmt.Println("Error, did not recognise number!")
		}
	}

	fmt.Println("Thanks for choosing our bank!")
}

func checkAccountBalance(balance float64) {
	fmt.Println("Your balance is: ", balance)
}

func bankOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check the balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}
