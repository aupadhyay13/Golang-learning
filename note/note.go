package note
import "time"
import "fmt"

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New(title, content string) (Note,error){
	if title == "" || content == ""{
		return Note{},errors.New("invalid Input")
	}
	retun Note{
		title : title,
		content: content,
		createdAt: time.Now(),
	},nil
}