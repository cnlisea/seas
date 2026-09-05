package code

import (
	"container/list"
)

type Code struct {
	Import   *list.List
	Main     *list.List
	MainLast *list.List
}

func New() *Code {
	return &Code{
		Import:   list.New(),
		Main:     list.New(),
		MainLast: list.New(),
	}
}

func (c *Code) Code() string {
	return ""
}
