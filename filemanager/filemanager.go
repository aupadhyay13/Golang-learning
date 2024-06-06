package filemanager
import "os"
import "bufio"
import "errors"
import "encoding/json"

type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func(fm FileManager) ReadLines() ([]string,error){
	file, err := os.Open(fm.InputFilePath)
	if err != nil{
		return nil,errors.New("failed to Open file")
	}
	scanner := bufio.NewScanner(file)
	
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil{
		file.Close()
		return nil,errors.New("failed to read line in file")
	}
	file.Close()
	return lines,nil
}

func(fm FileManager) WriteResult(data interface{}) error{
	file, err := os.Create(fm.OutputFilePath)
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

func New(inputPath, outputPath string) FileManager{
	return FileManager{
		InputFilePath : inputPath,
		OutputFilePath : outputPath,
	}
}