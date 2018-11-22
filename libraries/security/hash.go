package security

import (
	"crypto/sha1"
	"fmt"
)

func Hash(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}
