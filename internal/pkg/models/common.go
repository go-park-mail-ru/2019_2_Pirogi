package models

type ID int

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type Type struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

type Genre struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}

type Award struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}
