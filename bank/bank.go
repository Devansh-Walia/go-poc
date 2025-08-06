package bank

import (
	"fmt"
	"os"
	"strconv"

	investmentCalculator "github.com/investment-calculator/investment"
)

func writeToFile(balance float64) error {
	balanceText := fmt.Sprintf("%.2f", balance) // Format with 2 decimal places
	if err := os.WriteFile("passbook.txt", []byte(balanceText), 0644); err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return err
	}
	fmt.Println("Wrote to file!!")
	return nil
}

func loadDataFromFile() (float64, error) {
	data, err := os.ReadFile("passbook.txt")
	if err != nil {
		// Create file if it doesn't exist
		if os.IsNotExist(err) {
			fmt.Println("Creating new passbook file with initial balance 0")
			writeToFile(0)
			return 0, nil
		}
		return 0, fmt.Errorf("error reading passbook: %v", err)
	}

	if len(data) == 0 {
		fmt.Println("Passbook file is empty, starting with balance 0")
		writeToFile(0)
		return 0, nil
	}

	// Parse the text as float64
	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		fmt.Println("Passbook file has invalid data format, resetting to 0")
		writeToFile(0)
		return 0, nil
	}

	return balance, nil
}

func addToAccount(balance float64) error {
	var err error
	accountBalance, err = loadDataFromFile()
	if err != nil {
		return err
	}
	accountBalance += balance
	return writeToFile(accountBalance)
}

func withdrawFromAccount(balance float64) error {
	var err error
	accountBalance, err = loadDataFromFile()
	if err != nil {
		return err
	}

	if accountBalance >= balance {
		accountBalance -= balance
		return writeToFile(accountBalance)
	}
	return fmt.Errorf("insufficient balance: %.2f available, %.2f requested", accountBalance, balance)
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

var accountBalance float64

func Bank() {

	var err error
	accountBalance, err = loadDataFromFile()
	if err != nil {
		fmt.Printf("Error loading account balance: %v\n", err)
		return
	}
	fmt.Println("account balance", accountBalance)

menu:
	for {
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
			break menu
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
		fmt.Println("gb 👍")
	}

	writeToFile(accountBalance)

	fmt.Println("okay gg!")
}
