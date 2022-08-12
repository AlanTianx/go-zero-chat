package helper

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// 生成指定长度随机数 长度最小4 最大8
func RandNum(l int) int {
	if l < 4 || l > 8 {
		panic("长度最小4 最大8")
	}
	max := int(math.Pow(10, float64(l)) - 1)
	min := int(math.Pow(10, float64(l-1)))
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func RandStr(l int) (s string) {
	var chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	rand.Seed(time.Now().Unix())
	for i := 0; i < l; i++ {
		s = fmt.Sprintf("%s%s", s, string(chars[rand.Intn(len(chars))]))
	}
	return
}
