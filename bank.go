package main
import "fmt"
func main(){
	var accountBalance float64= 1000

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