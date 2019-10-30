package models

import "strconv"

type ID int

func (id ID) String() string {
	return strconv.Itoa(int(id))
}

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
