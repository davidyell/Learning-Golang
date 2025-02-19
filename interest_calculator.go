/*
An application written by Google Gemini in Go
This application tells you how much money you need in ordder to earn your desired sum at the set interst rate
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var interestRate float64
	var monthlyIncome float64
	var requiredPrincipal float64

	fmt.Print("Enter annual interest rate (as a percentage, e.g., 5): ")
	fmt.Scanln(&interestRate)
	interestRate /= 100 // Convert percentage to decimal

	fmt.Print("Enter desired monthly income: ")
	fmt.Scanln(&monthlyIncome)

	if interestRate <= 0 {
		fmt.Println("Interest rate must be positive.")
		return // Exit the program if the rate is invalid
	}

	monthlyInterestRate := interestRate / 12 // Annual to monthly
	requiredPrincipal = monthlyIncome / monthlyInterestRate

	fmt.Printf("Amount needed to generate $%.2f per month at %.2f%% interest: $%.0f\n", monthlyIncome, interestRate*100, math.Round(requiredPrincipal))

	// Example of compound interest calculation (optional)
	fmt.Println("\n--- Example of Compound Interest Calculation ---")
	var years int
	fmt.Print("Enter number of years to calculate compound interest: ")
	fmt.Scanln(&years)

	futureValue := requiredPrincipal * math.Pow(1+monthlyInterestRate, float64(years*12))
	fmt.Printf("Future value after %d years: $%.2f\n", years, futureValue)
}
