package filemanager
import "os"
import "fmt"
import "bufio"
import "errors"

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