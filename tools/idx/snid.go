package idx

import (
	"strconv"
	"time"
)

var (
	epoch     int64 = 1619827200000 // 设置雪花算法的起始时间戳，2021-05-01 00:00:00
	nodeBits  uint8 = 5             // 节点ID的位数
	stepBits  uint8 = 10            // 每毫秒序列号的位数
	nodeMax   int64 = -1 ^ (-1 << nodeBits)
	stepMax   int64 = -1 ^ (-1 << stepBits)
	timeShift uint8 = nodeBits + stepBits
	nodeShift uint8 = stepBits
)

type snowflake struct {
	lastTimestamp int64
	node          int64
	step          int64
}

func newSnowflake(node int64) *snowflake {
	return &snowflake{
		node: node,
	}
}

func (s *snowflake) next() int64 {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	if s.lastTimestamp == now {
		s.step++
		if s.step > stepMax {
			for now <= s.lastTimestamp {
				now = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		s.step = 0
	}
	s.lastTimestamp = now
	id := (now-epoch)<<timeShift | (s.node << nodeShift) | (s.step)
	return id
}

var (
	sf *snowflake
)

func init() {
	InitFlake(1)
}

func InitFlake(id int64) {
	sf = newSnowflake(id)
}

func GenSnId() int64 {
	return sf.next()
}

func GenSnIdStr() string {
	return strconv.FormatInt(GenSnId(), 10)
}
