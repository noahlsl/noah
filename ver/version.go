package ver

import (
	"fmt"
	"github.com/fatih/color"
)

func Version(args ...string) {
	_ = args
	fmt.Println("noah version", color.GreenString("1.0"))
}
