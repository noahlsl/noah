package idx

import (
	"github.com/google/uuid"
)

var (
	uu = uuid.New()
)

// GenUUID 生成UUID
func GenUUID() string {
	id := uu.String()
	return id
}
