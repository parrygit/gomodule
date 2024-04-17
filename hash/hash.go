package hash

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	re := h.Sum(nil)
	fmt.Println(re)
	md5Str := fmt.Sprintf("#{re}\n")
	return md5Str
}
