package models

import (
	"fmt"
	"strconv"
)

type ID int

func (id ID) String() string {
	return strconv.Itoa(int(id))
}

type FilmMark float32

func (fm FilmMark) String() string {
	return fmt.Sprintf("%f", fm)
}

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type Role struct {
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

type Like struct {
	ID       ID     `json:"id"`
	UserID   ID     `json:"user_id"`
	Target   string `json:"target"` // Film or person
	TargetID ID     `json:"target_id"`
}
