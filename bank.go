package main
import "fmt"
import "os"
import "strconv"
import "errors"


const accountBalanceFile = "balance.txt"
func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)   // generate a string 
	// []byte(balanceText)  converting string into bytes
	os.WriteFile(accountBalanceFile,[]byte(balanceText),0644) //0644 is a file permission to read and write file from owner
}

func getBalanceFromFile() (float64,error){
	// _ means we don't wanna work with it right now
	data, err := os.ReadFile(accountBalanceFile)
	// nil is a special value in go which stands for the absence
	if err != nil {	// Handling error
		return 1000, errors.New("failed To Find Balance File")
	}
	balanceText := string(data)
	balance , err := strconv.ParseFloat(balanceText, 64) // convert string to float
	if err != nil {			// if it can't convert string to float
		return 1000, errors.New("failed To Parse stored balance value")
	}
	return balance, nil
}

func main(){
	var accountBalance, err= getBalanceFromFile()
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


