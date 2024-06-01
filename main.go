//functions are nothing but value in golang

package main
import "fmt"

type funcType func(int) int

func main(){
	numbers := []int{1,2,3,4}
	moreNUmbers := []int{5,1,2}

	triple := tranformNumbers(&numbers, triple)
	doubled := tranformNumbers(&numbers, double)
	fmt.Println(doubled)
	fmt.Println(triple)

	tranformerFn1 := getTranformerFunction(&numbers)
	tranformerFn2 := getTranformerFunction(&moreNUmbers)

	tranformedNums := tranformNumbers(&numbers, tranformerFn1)
	moreTranformedNums := tranformNumbers(&moreNUmbers, tranformerFn2)

	fmt.Println(tranformedNums)
	fmt.Println(moreTranformedNums)
}

func getTranformerFunction(number *[]int) funcType{
	if((*number)[0] == 1){
		return double
	}else{
		return triple
	}
	
}

func tranformNumbers(numbers *[]int, tranform funcType) []int{
	dNumbers := []int{}
	for _,value := range *numbers{
		dNumbers = append(dNumbers,tranform(value))
	}

	return dNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}