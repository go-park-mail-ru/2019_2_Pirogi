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
	Status int    `json:"status" valid:"numeric"`
	Error  string `json:"error" valid:"alphanum"`
}

type Role string

type Target string

type Genre string

type Image struct {
	ID       ID     `json:"id" valid:"numeric"`
	Filename string `json:"filename" valid:"image"`
}

type ReviewsNum int
