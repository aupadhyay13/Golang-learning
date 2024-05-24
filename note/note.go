package note
import "time"
import "fmt"
import "errors"
import "os"
import "strings"
import "encoding/json"

type Note struct {
	title string
	content string
	createdAt time.Time
}
func (note Note) DisplayNote() {
	fmt.Printf("Your Note Titled %v has the following content: \n\n%v\n\n ",note.title, note.content)
} 

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName,json,0644)
	
}

func New(title, content string) (Note,error){
	if title == "" || content == ""{
		return Note{},errors.New("invalid Input")
	}
	return Note{
		title : title,
		content: content,
		createdAt: time.Now(),
	},nil
}