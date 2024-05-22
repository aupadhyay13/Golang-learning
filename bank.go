package main

import "fmt"
import "errors"
import "os"

// Goals

// 1) Validate User Input
// 	=> Show error message and exit if Invalid input is provided
// 	- No Negative Numbers
// 	- Not 0
// 2) Store Calculated results into file

func validateUserInput(revenue,expenses,taxRate float64) error {
	if revenue <= 0 || expenses <= 0 || taxRate <= 0 {
		return errors.New("invalid Input Provided! No Negative Number")
	}
	return nil
}	

func writeResultIntoFile(ebt,profit,ratio float64){
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n",ebt,profit,ratio) //generate string with formatting 
	os.WriteFile("result.txt",[]byte(results),0644)
}
func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")
	err := validateUserInput(revenue,expenses,taxRate)
	if err != nil{
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeResultIntoFile(ebt, profit, ratio)
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}