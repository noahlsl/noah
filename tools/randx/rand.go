package randx

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/noahlsl/noah/consts"
)

// GetRand 随机字符串
func getRand(size int, kind int) string {
	iKind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if isAll { // random iKind
			iKind = rand.Intn(3)
		}
		scope, base := kinds[iKind][0], kinds[iKind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// GetRandNum 随机数字
func GetRandNum(size int, f ...int) int64 {

	var r string
	for {
		r = getRand(size, consts.KcRandKindNum)
		if len(f) == 0 {
			break
		}
		split := strings.Split(r, "")
		if split[0] != "0" {
			break
		}
	}
	r = getRand(size, consts.KcRandKindNum)
	i, _ := strconv.ParseInt(r, 10, 64)
	return i
}

// GetRandLowerStr 随机小写字母
func GetRandLowerStr(size int) string {
	return getRand(size, consts.KcRandKindLower)
}

// GetRandUpperStr 随机大写字母
func GetRandUpperStr(size int) string {
	return getRand(size, consts.KcRandKindUpper)
}

// GetRandStr 随机字符串
func GetRandStr(size int) string {
	return getRand(size, consts.KcRandKindAll)
}
