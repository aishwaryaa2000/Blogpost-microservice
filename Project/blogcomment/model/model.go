package model
import(
	"github.com/gofrs/uuid"
)

type Comment struct{
	CommentId uuid.UUID `json:"commentId"`
	Message string `json:"message"`
}

type EventComment struct{
	Type string `json:"type"`
	Data map[string]interface{} 
}


