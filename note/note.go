package note
import "time"
import "fmt"
import "errors"
import "os"
import "strings"
import "encoding/json"

type Note struct {
	Title string
	Content string
	CreatedAt time.Time
}
func (note Note) DisplayNote() {
	fmt.Printf("Your Note Titled %v has the following content: \n\n%v\n\n ",note.Title, note.Content)
} 

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
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
		Title : title,
		Content: content,
		CreatedAt: time.Now(),
	},nil
}