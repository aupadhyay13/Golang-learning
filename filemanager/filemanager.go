package filemanager
import "os"
import "bufio"
import "errors"
import "encoding/json"

func ReadLines(path string) ([]string,error){
	file, err := os.Open(path)
	if err != nil{
		return nil,errors.New("Failed to Open file!")
	}
	scanner := bufio.NewScanner(file)
	
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil{
		file.Close()
		return nil,errors.New("Failed to read line in file!")
	}
	file.Close()
	return lines,nil
}

func WriteJSON(path string,data interface{}) error{
	file, err := os.Create(path)
	if err != nil{
		return errors.New("failed To create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil{
		return errors.New("failed to convert json data")
	}
	file.Close()
	return nil
}