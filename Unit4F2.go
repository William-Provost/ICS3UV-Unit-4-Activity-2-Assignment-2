// Author: William Provost
// Version: 1.0.0
// Date: 2025-11-28
// Fileoverview: This program calculates how many years it will take 
// for a bank account to reach a target amount based on yearly interest

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variable declaration
	var startingAmount float64
	var interestRate float64
	var targetAmount float64
	var currentAmount float64
	var years int

	reader := bufio.NewReader(os.Stdin)

	// get user input
	fmt.Print("Enter the starting bank account amount: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	startingAmount, _ = strconv.ParseFloat(input, 64)

	fmt.Print("Enter the yearly interest rate (as a percentage): ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	interestRate, _ = strconv.ParseFloat(input, 64)

	fmt.Print("Enter the amount needed for post-secondary education: ")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	targetAmount, _ = strconv.ParseFloat(input, 64)

	// initialize
	currentAmount = startingAmount
	years = 0

	// calculate number of years to reach target
	for currentAmount < targetAmount {
		currentAmount += currentAmount * (interestRate / 100)
		years++
	}

	// output result
	fmt.Printf("\nIt will take %d years for the starting bank account to reach the required amount for post-secondary education.\n", years)

	fmt.Println("\nDone.")
}
