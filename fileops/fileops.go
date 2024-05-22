package fileops

import ("fmt"
"os"
"errors"
"strconv"
)
func writeFloatToFile(fileName string, value float64) {
	valueText := fmt.Sprint(value)   // generate a string 
	// []byte(balanceText)  converting string into bytes
	os.WriteFile(fileName,[]byte(valueText),0644) //0644 is a file permission to read and write file from owner
}

func getFloatFromFile(fileName string) (float64,error){
	// _ means we don't wanna work with it right now
	data, err := os.ReadFile(fileName)
	// nil is a special value in go which stands for the absence
	if err != nil {	// Handling error
		return 1000, errors.New("failed To Find File")
	}
	valueText := string(data)
	value , err := strconv.ParseFloat(valueText, 64) // convert string to float
	if err != nil {			// if it can't convert string to float
		return 1000, errors.New("failed To Parse stored balance value")
	}
	return value, nil
}