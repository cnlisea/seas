package code

func (c *Code) MainWriteString(s ...string) {
	for i := range s {
		c.Main.PushBack(s[i])
	}
}
