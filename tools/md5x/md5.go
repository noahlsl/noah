package md5x

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"github.com/noahlsl/noah/tools/strx"
)

/** 加密方式 **/

func ByString(str string, salts ...interface{}) string {
	return ByBytes(strx.S2b(str), salts...)
}

func ByBytes(b []byte, salts ...interface{}) string {

	var s []byte
	for _, salt := range salts {
		s = append(s, []byte(fmt.Sprintf("%v", salt))...)
	}
	h := md5.New()
	h.Write(s) // 先写盐值
	h.Write(b)

	return hex.EncodeToString(h.Sum(nil))
}
