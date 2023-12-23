package sess

import (
	"crypto/rand"
	"fmt"
)

func GenerateUUID() string {
	buf := make([]byte, 16)
	 rand.Read(buf)
	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
