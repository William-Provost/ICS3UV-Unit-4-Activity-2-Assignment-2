/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-28
 * @fileoverview This program calculates how many years it will take 
 * for a bank account to reach a target amount based on yearly interest.
 */

// variables
let startingAmount: number = 0;
let interestRate: number = 0;
let targetAmount: number = 0;
let currentAmount: number = 0;
let years: number = 0;

// get user input
startingAmount = parseFloat(prompt("Enter the starting bank account amount: ") || "0");
interestRate = parseFloat(prompt("Enter the yearly interest rate (as a percentage): ") || "0");
targetAmount = parseFloat(prompt("Enter the amount needed for post-secondary education: ") || "0");

// initialize
currentAmount = startingAmount;
years = 0;

// calculate number of years to reach target
while (currentAmount < targetAmount) {
  currentAmount += currentAmount * (interestRate / 100);
  years++;
}

// output result
console.log(`\nIt will take ${years} years for the starting bank account to reach the required amount for post-secondary education.`);

console.log("\nDone.");
