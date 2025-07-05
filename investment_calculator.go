package main

import (
	"fmt"
	"math"
)

func main() {

	const inflationRate float64 = 2.0

	var principalAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Println("Investment Calculator")
	fmt.Print("Enter the principal amount: ")
	fmt.Scan(&principalAmount)
	fmt.Print("Enter the expected annual return rate (in %): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter the number of years to invest: ")
	fmt.Scan(&years)

	futureValue := principalAmount * math.Pow(float64(1+expectedReturnRate/100), float64(years))
	futureAdjustedValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Printf("With a principal of: %v\n with an anual rate of: %v\n and a investment horizon of %2f\n Future value: %.2f\n Future adjusted value: %.2f\n", principalAmount, expectedReturnRate, years, futureValue, futureAdjustedValue)

}
