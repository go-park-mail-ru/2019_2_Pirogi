package models

import (
	"fmt"
	"strconv"
)

type ID int

func (id ID) String() string {
	return strconv.Itoa(int(id))
}

// number of stars
type Mark float32

func (fm Mark) String() string {
	return fmt.Sprintf("%f", fm)
}

type Error struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

type Role string

type Target string

type Genre string

type Image struct {
	ID       ID     `json:"id"`
	Filename string `json:"filename"`
}

/*
type Award struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
}*/
