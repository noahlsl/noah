package hash

import (
	"strconv"
	"testing"
)

func TestGetOneHash(t *testing.T) {
	var users []string
	for i := 0; i < 10; i++ {
		users = append(users, strconv.Itoa(i))
	}

	for _, u := range users {
		t.Log(GetOneHash(u))
	}
}
