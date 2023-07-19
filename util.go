package sess

import (
	"crypto/rand"
	"fmt"
)

func GenerateUUID() string {
	runes := []rune("0123456789abcdef")

	buf := make([]byte, 32)
	rand.Read(buf)

	result := make([]rune, 32)
	for k, v := range buf {
		result[k] = runes[int(v)%16]
	}

	p1 := result[0:8]
	p2 := result[8:12]
	p3 := result[12:16]
	p4 := result[16:20]
	p5 := result[20:32]

	return fmt.Sprintf("%s-%s-%s-%s-%s", string(p1), string(p2), string(p3), string(p4), string(p5))
}
