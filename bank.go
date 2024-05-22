package main
import "fmt"


const accountBalanceFile = "balance.txt"


func main(){
	var accountBalance, err= getFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		panic("Can't continue Sorry!") // built in Panic function help in stop the execution of program
		// return // if don't want to continue execution if there is error

	}
	fmt.Println("Welcome to Go Bank!!")

   for { 	

	presentOptions()

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
			writeFloatToFile(accountBalanceFile,accountBalance)
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


