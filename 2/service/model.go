package service

// create a structur
// library
// id
// name
// author
// year
// isbn
// pages
// publisher
// language
// description
// cover
// file
// category
// status
// created_at
// updated_at
// deleted_at

type Library struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Author      string `json:"author"`
	Year        int    `json:"year"`
	ISBN        string `json:"isbn"`
	Pages       int    `json:"pages"`
	Publisher   string `json:"publisher"`
	Language    string `json:"language"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
	File        string `json:"file"`
	Category    string `json:"category"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	DeletedAt   string `json:"deleted_at"`
}
