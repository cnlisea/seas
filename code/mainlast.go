package code

func (c *Code) MainLastWriteString(s ...string) {
	for i := range s {
		c.MainLast.PushBack(s[i])
	}
}
