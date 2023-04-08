package idx

import (
	"strconv"
	"time"

	"github.com/noahlsl/noah/tools/randx"
)

// GenOid 生成单号
// 比如：单号-DW2023022812121265845
// DW为前缀,中间是单号的年月日时分秒，最后5位是随机数字
func GenOid(snPrefix string) string {
	return snPrefix + time.Now().Format("20060102150405") + strconv.Itoa(int(randx.GetRandNum(5)))
}

// GenDid 生成日期单号
// 比如：单号-DW20230228
// DW为前缀,中间是单号的年月日
func GenDid(snPrefix string) string {
	return snPrefix + time.Now().Format("20060102")
}

// GenDTid 生成日期时间单号
// 比如：单号-DW20230228
// DW为前缀,中间是单号的年月日
func GenDTid(snPrefix string) string {
	return snPrefix + time.Now().Format("20060102150405")
}
