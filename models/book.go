package models

type Book struct {
	BookID uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`  
	Author string `json:"author"`
	Description   string `json:"Description"`  
}