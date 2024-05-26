package main
import "fmt"

func main(){
	var productNames [4]string = [4]string{"A book"}
	prices := [4]float64{10.99,9.99,45.99,20.00}
	fmt.Println(prices)

	productNames[2] = "A Carpet"
	fmt.Println(productNames)

	fmt.Println(prices[2])
	// featuredPrices := prices[:3]  // from begineeing
	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99 // though we changed in featuredPrice it will change in price array as well
	highlitedPrices := featuredPrices[:1]
	fmt.Println(highlitedPrices)

	fmt.Println(featuredPrices)
	fmt.Println(prices)
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))
}