package entity

import (
	"fmt"
	"strconv"
)

type ID int

//Album data
type Album struct {
	Title  string `redis:"title"`
	Artist string `redis:"artist"`
	ID     ID     `redis:"id"`
}

func StringToID(s string) ID {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("invalid input to ID: %q", s))
	}
	return ID(i)
}

func (i ID) String() string {
	return strconv.Itoa(int(i))
}
