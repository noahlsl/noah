package logx

import (
	"fmt"
	"github.com/noahlsl/noah/tools/structx"
	"github.com/noahlsl/noah/tools/strx"
	"github.com/thinkeridea/go-extend/exbytes"
	"github.com/thinkeridea/go-extend/exstrings"
	"os"
	"strings"
)

type stackItem struct {
	Func   string `json:"func"`
	Line   string `json:"line"`
	Source string `json:"source"`
}

func GetStack(err error) string {

	s := washPath(fmt.Sprintf("%+v", err))
	s1 := strings.Split(s, "\n")
	if len(s1) == 0 {
		return ""

	} else if len(s1) == 1 {
		return s1[0]
	}

	var items []stackItem
	for i := 1; i < len(s1); i += 2 {
		item := stackItem{}
		f := strings.Split(s1[i], "/")
		f1 := strings.Split(f[len(f)-1], ".")
		if len(f1) == 2 {
			// 过滤外部包
			if strings.HasPrefix(f1[0], "go.") ||
				filter(f1[0]) {
				continue
			}

			item.Source = f1[0]
			item.Func = f1[1]
		}
		item.Line = strings.TrimSpace(s1[i+1])
		items = append(items, item)

	}
	if len(items) == 0 {
		return ""
	}

	return strx.B2s(structx.StructToBytes(items))
}

// 路径脱敏
func washPath(s string) string {
	sb := exstrings.Bytes(s)
	path, _ := os.Getwd()
	pathByte := exbytes.Replace([]byte(path+"/"), []byte("\\"), []byte("/"), -1)
	root := os.Getenv("GOROOT")
	rootByte := exbytes.Replace([]byte(root+"/"), []byte("\\"), []byte("/"), -1)
	sb = exbytes.Replace(sb, pathByte, []byte(""), -1)
	sb = exbytes.Replace(sb, rootByte, []byte(""), -1)
	return exbytes.ToString(sb)
}
