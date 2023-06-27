package common

import (
	"math/rand"
	"strings"
	"time"
)

// GenerateRandomString 生成包含 1-9、a-z 的指定长度的随机字符串
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())

	var builder strings.Builder
	chars := "0123456789abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < length; i++ {
		builder.WriteByte(chars[rand.Intn(len(chars))])
	}
	return builder.String()
}
