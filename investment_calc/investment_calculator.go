package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 6.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// Formatting
	//formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	//formattedRFV := fmt.Sprintf("Future Value adjusted for inflation: %.1f\n", futureRealValue)
	//fmt.Print(formattedFV, formattedRFV)

	//fmt.Printf("Future Value: %.1f\nFuture Value adjusted for inflation: %.1f\n", futureValue, futureRealValue)

	fmt.Printf(`Future Value: %.1f
	Future Value adjusted for inflation: %.1f`,
		futureValue, futureRealValue)
}
