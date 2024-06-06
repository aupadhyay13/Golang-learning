package prices
import "fmt"
import "os"
import "bufio"

import "example.com/struct-project/conversion"

type TaxIncludedPriceJob struct{
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]float64
}

func(job *TaxIncludedPriceJob) LoadData(){
	file, err := os.Open("prices.txt")
	if err != nil{
		fmt.Println("An error Occurred!")
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(file)
	
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil{
		fmt.Println("An error Occurred!")
		fmt.Println(err)
		file.Close()
		return
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}
	job.InputPrices = prices
	file.Close()
}


func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)
	for _,price := range job.InputPrices {

		TaxIncludedPrice := price * (1+job.TaxRate)
		result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",TaxIncludedPrice)
	}
	fmt.Println(result)
}



func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}