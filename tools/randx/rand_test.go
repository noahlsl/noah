package randx

import (
	"testing"
)

func TestGetRandLowerStr(t *testing.T) {
	t.Log(GetRandLowerStr(5))
}

func TestGetRandNum(t *testing.T) {
	t.Log(GetRandNum(5))
}

func TestGetRandUpperStr(t *testing.T) {
	t.Log(GetRandUpperStr(5))
}

func TestGetRandStr(t *testing.T) {
	t.Log(GetRandStr(5))
}
