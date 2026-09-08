package code

import (
	"container/list"
	"strings"
)

type Code struct {
	Name     string
	Version  string
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

func (c *Code) Init() error {
	c.Main.PushBack("func main() {\n")
	c.Main.PushBack("\tvar (\n")
	c.Main.PushBack("\t\ta = app.New()\n")
	c.Main.PushBack("\t\terr   error\n")
	c.Main.PushBack("\t)\n")
	// main last
	c.MainLast.PushBack("\tif err = a.Run(); err != nil {\n")
	c.MainLast.PushBack("\t\tpanic(err)\n")
	c.MainLast.PushBack("\t}\n")
	c.MainLast.PushBack("}\n")
	return nil
}

func (c *Code) Code() string {
	var buffer strings.Builder
	buffer.WriteString("package main\n\n")
	buffer.WriteString("import (\n")
	buffer.WriteString("\t\"github.com/cnlisea/ant/app\"\n")

	var e *list.Element
	if c.Import.Len() > 0 {
		for e = c.Import.Front(); e != nil; e = e.Next() {
			buffer.WriteString(e.Value.(string))
		}
	}
	buffer.WriteString(")\n\n")

	for e = c.Main.Front(); e != nil; e = e.Next() {
		buffer.WriteString(e.Value.(string))
	}
	for e = c.MainLast.Front(); e != nil; e = e.Next() {
		buffer.WriteString(e.Value.(string))
	}
	return buffer.String()
}
