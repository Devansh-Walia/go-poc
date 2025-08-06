package main

import (
	"fmt"

	investmentCalculator "github.com/investment-calculator/investment"
)

var account_balance float64 = 0

func add_to_account(balance float64) {
	account_balance += balance
}

func withdraw_from_account(balance float64) {
	if account_balance > balance {
		account_balance -= balance
	} else {
		fmt.Println("Insufficient balance")
	}
}

func GetVar(message string) (value float64) {
	fmt.Print(message)
	fmt.Scan(&value)
	return value
}

func investment() {
	var investmentAmount, years, investmentRate float64

	investmentAmount = GetVar("Enter the rate you want to invest: ")
	years = GetVar("Enter the no of years you want to invest: ")
	investmentRate = GetVar("enter the rate of investment: ")

	futureValue, inflationAdjustedFutureValue, profit := investmentCalculator.Calculation(investmentAmount, years, investmentRate)

	fmt.Println("future value -> ", futureValue)
	fmt.Println("inflation adjusted future value ->", inflationAdjustedFutureValue)

	fmt.Printf("Total profit of %.2f was achieved", profit)
}

func main() {
	fmt.Println("Please choose from the following options:")
	fmt.Println("To Add money -> 1")
	fmt.Println("To Invest -> 2")
	fmt.Println("To Withdraw -> 3")
	fmt.Println("To Exit -> 4")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		fmt.Println("Add money")
		add_to_account(GetVar("Enter the amount you want to add: "))
	case 2:
		fmt.Println("Invest")
		investment()
	case 3:
		fmt.Println("Withdraw")
		withdraw_from_account(GetVar("Enter the amount you want to withdraw: "))
	case 4:
		fmt.Println("Exit")
	default:
		fmt.Println("Invalid choice")
	}
}
