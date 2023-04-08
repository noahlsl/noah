package help

import (
	"fmt"
	"github.com/fatih/color"
)

var k = "    "

func Help(args ...string) {
	_ = args
	fmt.Println("This is an Noah CMD tool")

	color.Blue("Usage:")
	color.Green(k + "noah [command]")
	fmt.Println()

	// Command命令说明
	color.Blue("Available Commands:")
	fmt.Println(k, color.YellowString(getStr("new")), k, "Create a new project")
	fmt.Println(k, color.CyanString(getStr("model")), k, "Generate Model files based on SQL")
	fmt.Println(k, color.YellowString(getStr("help")), k, "Help about any command")
	fmt.Println()

	// Flags说明
	color.Blue("Flags:")
	color.Green(k + "-h, --help      help for noah")
	color.Green(k + "-v, --version   version for noah\n\n")

	fmt.Println("Use " + color.GreenString("noah") + " [command] --help for more information about a command.")
}

func getStr(s string) string {
	l := 10 - len(s)
	for i := 0; i < l; i++ {
		s += " "
	}
	return s
}
