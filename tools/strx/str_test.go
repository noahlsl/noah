package strx

import "testing"

func TestB2s(t *testing.T) {
	B2s([]byte("hello"))
}

func TestS2b(t *testing.T) {
	S2b("hello")
}
