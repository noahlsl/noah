package main

import (
	"os"

	"github.com/noahlsl/noah/help"
	"github.com/noahlsl/noah/newx"
	"github.com/noahlsl/noah/ver"
)

var cmdMap map[string]func(args ...string)

func init() {
	cmdMap = make(map[string]func(args ...string))
	cmdMap["-h"] = help.Help
	cmdMap["--help"] = help.Help
	cmdMap["-v"] = ver.Version
	cmdMap["--version"] = ver.Version
	cmdMap["new"] = newx.New
}

func main() {
	args := os.Args
	if len(args) == 1 {
		help.Help()
		return
	}

	if fn, ok := cmdMap[args[1]]; ok {
		fn(args[2:]...)
		return
	}

	help.Help()
}
