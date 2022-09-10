package model
import(
	"github.com/gofrs/uuid"
)

type EventPost struct{
	Type string `json:"type"`
	Data map[string]interface{} 
}

type Post struct{
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
}