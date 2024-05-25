package todo
import "fmt"
import "errors"
import "os"
import "encoding/json"

type Todo struct {
	Text string `json:"text"`
}
func (todo Todo) DisplayToDo() {
	fmt.Println(todo.Text)
} 

func (todo Todo) Save() error {
	fileName := "todo.json"
	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName,json,0644)
	
}

func New(content string) (Todo,error){
	if content == ""{
		return Todo{},errors.New("invalid Input")
	}
	return Todo{
		Text: content,
	},nil
}