package logx

import "strings"

// 过滤的字段
var filterPath = []string{
	"runtime",
	"github",
	"gitee",
	"gitlab",
	"gopkg",
	"middleware",
	"google",
}

func filter(f string) bool {
	for _, path := range filterPath {
		if strings.Contains(f, path) {
			return true
		}
	}
	return false
}
