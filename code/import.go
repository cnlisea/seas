package code

func (c *Code) ImportWriteString(s ...string) {
	for i := range s {
		c.Import.PushBack(s[i])
	}
}
