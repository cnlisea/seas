package code

import (
	"container/list"
)

func (c *Code) Clone() *Code {
	var (
		clone = New()
		e     *list.Element
	)
	for e = c.Import.Front(); e != nil; e = e.Next() {
		clone.Import.PushBack(e.Value.(string))
	}
	for e = c.Main.Front(); e != nil; e = e.Next() {
		clone.Main.PushBack(e.Value.(string))
	}
	for e = c.MainLast.Front(); e != nil; e = e.Next() {
		clone.MainLast.PushBack(e.Value.(string))
	}
	return clone
}
