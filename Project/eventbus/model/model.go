//This package consists of the structure declarations
package model

type Event struct{
	Type string `json:"type"`
	Data map[string]interface{} `json:"Data"`
}


/*
	Here,the Event struct can be either of comment or post

	If Event has the type as comment created
	then data interface will have commentId, message, id
	Example JSON - 
	{
		"type" : "comment created"
		"data" : {
					"id" : post ID here
					"commentId" : comment ID here
					"message" : comment msg here
				 }
	}

	If Event has the type as Post created
	then data interface will have id and title
	Example JSON - 
	{
		"type" : "Post created"
		"data" : {
					"id" : post ID here
					"title" : post title here
				 }
	}
*/