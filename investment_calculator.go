package main

import (
	"fmt"
	"math"
)

func main() {
	investmentAmount, years, investmentRate := 10000.0, 10.0, 5.5

	futureValue := investmentAmount * math.Pow((1+investmentRate/100), years)

	fmt.Println(futureValue)
}
