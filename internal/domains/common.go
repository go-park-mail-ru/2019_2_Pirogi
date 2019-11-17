package domains

import (
	"fmt"
	"strconv"
)

type ID int

func (id *ID) String() string {
	return strconv.Itoa(int(*id))
}

type Mark float32

func (fm Mark) String() string {
	return fmt.Sprintf("%f", fm)
}

type Error struct {
	Status int    `json:"status" valid:"numeric"`
	Error  string `json:"error" valid:"optional"`
}

type Role string

type Target string

type Genre string

type Image string

type TrailerWithTitle struct {
	Title   string `json:"title" valid:"text"`
	Trailer string `json:"trailer" valid:"link"`
}
