package domains

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/pkg/hash"
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

type Role string

type Target string

type Genre string

type Image string

type TrailerWithTitle struct {
	Title   string `json:"title" valid:"text"`
	Trailer string `json:"trailer" valid:"link"`
}

func (i *Image) Hash() {
	*i = Image(hash.SHA1(string(*i)))
}
