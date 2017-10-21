package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)

// 根据两个key，生成一个长度为10的令牌
func NewToken(k1, k2 string) string {
	t := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(t, 10)+k1+k2)
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token[0:11]
}
