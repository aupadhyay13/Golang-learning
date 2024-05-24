package main
import "fmt"

type str string

func (text str) log(){		// added methods to custom types
	fmt.Println(text)
}

func main(){
	var name str = "Aditya"

	name.log()
}