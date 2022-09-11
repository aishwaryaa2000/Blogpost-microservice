//This package consists of the structure declarations
package model
import(
	"github.com/gofrs/uuid"
)

type EventPost struct{
	Type string `json:"type"`
	Data map[string]interface{} 
}

/*
	Here,the EventPost struct has the type as Post created
	and data interface will have id and title
	Example JSON - 
	{
		"type" : "Post created"
		"data" : {
					"id" : post ID here
					"title" : post title here
				 }
	}
*/

type Post struct{
	Id uuid.UUID `json:"id"`
	Title string `json:"title"`
}
