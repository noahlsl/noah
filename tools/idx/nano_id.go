package idx

import "github.com/matoous/go-nanoid/v2"

// GetNanoId 获取NanoId
// 该ID替代UUID,占用的字节数更小,生成速度更快
func GetNanoId() string {
	id, _ := gonanoid.New()
	return id
}
