package tool

import (
	"crypto/md5"
	"fmt"
)

func Md5ByString(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
