package investmentCalculator

import (
	"fmt"
	"math"
)

const InflationRate = 4.5

func InvestmentCalculator() {
	var investmentAmount, years, investmentRate float64

	investmentAmount = GetVar("Enter the rate you want to invest: ")
	years = GetVar("Enter the no of years you want to invest: ")
	investmentRate = GetVar("enter the rate of investment: ")

	futureValue, inflationAdjustedFutureValue, profit := Calculation(investmentAmount, years, investmentRate)

	fmt.Println("future value -> ", futureValue)
	fmt.Println("inflation adjusted future value ->", inflationAdjustedFutureValue)

	fmt.Printf("Total profit of %.2f was achieved", profit)
}

func Calculation(investmentAmount, years, investmentRate float64) (fv float64, inv float64, profit float64) {
	fv = investmentAmount * math.Pow((1+investmentRate/100), years)
	inv = fv / math.Pow(InflationRate/100, years)
	profit = inv - investmentAmount

	// alternative  to return fv, inv, profit
	return
}

func GetVar(message string) (value float64) {
	fmt.Print(message)
	fmt.Scan(&value)

	return value
}
