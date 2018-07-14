package ut

import (
	"crypto/md5"
	"fmt"
)

// Md5 returns the MD5 checksum of the data.
func Md5(text string) string {
	sign := md5.New()
	sign.Write([]byte(text))
	return fmt.Sprintf("%x", sign.Sum(nil))
}
