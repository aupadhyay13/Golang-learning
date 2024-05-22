package main
import "fmt"
import "os"
import "strconv"


const accountBalanceFile = "balance.txt"
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)   // generate a string 
	// []byte(balanceText)  converting string into bytes
	os.WriteFile(accountBalanceFile,[]byte(balanceText),0644) //0644 is a file permission to read and write file from owner
}

func getBalanceFromFile() float64{
	// _ means we don't wanna work with it right now
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance , _ := strconv.ParseFloat(balanceText, 64) // convert string to float
	return balance
}

func main(){
	var accountBalance float64= getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!!")

   for { 	

	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	switch choice {
		case 1:
			fmt.Println("Your Balance is: ", accountBalance)
		case 2:
			fmt.Println("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount < 0 {
				fmt.Println("Invalid DepositAmount! Must be greater than 0!")
				// return // stops the execution of function
				continue //skip the current loop and iterate next
			}

			accountBalance += depositAmount
			fmt.Println("Balance Updated! New Amount: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3: 
			fmt.Println("Your Amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount < 0 {
				fmt.Println("Invalid DepositAmount! Must be greater than 0!")
				// return // stops the execution of function
				continue 
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Invalid Amount! You can't withdraw more tha what you have!")
				// return
				continue 
			}

			accountBalance -= withdrawAmount
			fmt.Println("Balance Updated! New Amount: ",accountBalance)
		default:
			fmt.Println("GoodBye!!!")
			return
			//break

	}

}
	fmt.Println("Thanks for choosing our bank!")
   
}