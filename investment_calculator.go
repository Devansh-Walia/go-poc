package main

import (
	"fmt"
	"math"
)

const inflationRate = 4.5

func main() {
	var investmentAmount, years, investmentRate float64

	investmentAmount = getVar("Enter the rate you want to invest: ")
	years = getVar("Enter the no of years you want to invest: ")
	investmentRate = getVar("enter the rate of investment: ")

	futureValue, inflationAdjustedFutureValue, profit := calculation(investmentAmount, years, investmentRate)

	fmt.Println("future value -> ", futureValue)
	fmt.Println("inflation adjusted future value ->", inflationAdjustedFutureValue)

	fmt.Printf("Total profit of %.2f was achieved", profit)
}

func calculation(investmentAmount, years, investmentRate float64) (fv float64, inv float64, profit float64) {
	fv = investmentAmount * math.Pow((1+investmentRate/100), years)
	inv = fv / math.Pow(inflationRate/100, years)
	profit = inv - investmentAmount

	// alternative  to return fv, inv, profit
	return
}

func getVar(message string) (value float64) {
	fmt.Print(message)
	fmt.Scan(&value)

	return value
}
