package main
import "fmt"
import "errors"

func main(){
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
}



func getNoteData() (string,string,error){
	title,err := getUserInput("Note Title:")
	if err != nil{
		return "","",err
	}
	content,err := getUserInput("Note Content:")
	if err != nil{
		return "","",err
	}
	return title,content,nil

}

func getUserInput(prompt string) (string,error){
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "",errors.New("invalid Input")
	}
	return value,nil
}