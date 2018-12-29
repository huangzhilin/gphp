package gphp

import (
	"crypto/md5"
	"fmt"
)

//md5加密
func Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
