package aesx

import "testing"

func TestAes(t *testing.T) {
	c := Cfg{
		Key: "0123456789abcdef",
		IV:  "0123456789abcdef",
	}
	cipher := c.NewCipher()
	str := EncodeToStr(cipher, "hello world")
	toStr := DecodeToStr(cipher, str)
	t.Log(toStr)
}
