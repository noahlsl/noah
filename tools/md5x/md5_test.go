package md5x

import (
	"testing"

	"github.com/noahlsl/noah/tools/strx"
)

func TestByString(t *testing.T) {
	byString := ByString("test", 1, 2)
	t.Log(byString)
}

func TestByBytes(t *testing.T) {
	by := ByBytes(strx.S2b("test"), 1, 2)
	t.Log(by)
}
