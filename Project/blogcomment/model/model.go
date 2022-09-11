//This package consists of the structure declarations
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

/*
	Here,the EventComment struct has the type as comment created
	and data interface will have commentId, message, id
	Example JSON - 
	{
		"type" : "comment created"
		"data" : {
					"id" : post ID here
					"commentId" : comment ID here
					"message" : comment msg here
				 }
	}
*/