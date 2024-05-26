package main
import "fmt"

func main(){
	prices := []float64{10.99,8.99}	
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)
	fmt.Println(prices)
}