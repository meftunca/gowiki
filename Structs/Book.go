package Structs

type POSTBookStruct struct {
	Name        string `json:"name"`
	AuthorName  string `json:"authorName"`
	Publisher   string `json:"publisher"` //yayın evi
	PageCount   int    `json:"pageCount"`
	description string `json:"description"`
	language    int    `json:"language"`
	ISBN        string `json:"isbn"`
	ReleaseDate int    `json:"releaseDate"`
}

type PUTBookStruct struct {
	Name        string `json:"name"`
	AuthorName  string `json:"authorName"`
	Publisher   string `json:"publisher"` //yayın evi
	PageCount   int    `json:"pageCount"`
	description string `json:"description"`
	language    int    `json:"language"`
	ISBN        string `json:"isbn"`
	ReleaseDate int    `json:"releaseDate"`
}

type GETBookStruct struct {
	id          int    `json:"id"`
	Name        string `json:"name"`
	AuthorName  string `json:"authorName"`
	Publisher   string `json:"publisher"` //yayın evi
	PageCount   int    `json:"pageCount"`
	description string `json:"description"`
	language    int    `json:"language"`
	ISBN        string `json:"isbn"`
	ReleaseDate int    `json:"releaseDate"`
}

type DELETEBookStruct struct {
	id int `json:"id"`
}
