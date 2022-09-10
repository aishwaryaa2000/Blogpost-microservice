package model

import(
	"github.com/gofrs/uuid"
)

type SinglePost struct{
	Id uuid.UUID `json:"id"`
	Comments map[uuid.UUID]string `json:"comments"`
	Title string `json:"title"`
}

type Event struct{
	Type string `json:"type"`
	Data map[string]interface{} `json:"Data"`
}
