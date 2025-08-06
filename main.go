package main

import (
	"fmt"

	investmentCalculator "github.com/investment-calculator/investment"
)

var accountBalance float64 = 0

func addToAccount(balance float64) {
	accountBalance += balance
}

func withdrawFromAccount(balance float64) {
	if accountBalance > balance {
		accountBalance -= balance
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

	for i := 0; i < 2; i++ {
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
			addToAccount(GetVar("Enter the amount you want to add: "))
		case 2:
			fmt.Println("Invest")
			investment()
		case 3:
			fmt.Println("Withdraw")
			withdrawFromAccount(GetVar("Enter the amount you want to withdraw: "))
		case 4:
			fmt.Println("Exit")
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
		fmt.Println("gb 👍")
	}
}
