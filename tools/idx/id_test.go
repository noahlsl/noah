package idx

import (
	"fmt"
	"testing"
)

func TestGetNanoId(t *testing.T) {
	fmt.Println(GetNanoId())
}

func TestGenOid(t *testing.T) {
	fmt.Println(GenOid("DW"))
}

func TestGenSnId(t *testing.T) {
	id := GenSnIdStr()
	fmt.Println(id)
	fmt.Println(len(id))
}
