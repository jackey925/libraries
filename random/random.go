package random

import (
	"math/rand"
	"time"
)

// RandInt32 随机生成一个整数切片
func RandInt32(n int32) []int32 {
	rand.Seed(time.Now().Unix())
	var list []int32
	for i := 0; i < int(n); i++ {

		list = append(list, int32(rand.Intn(int(n))))
	}
	return list
}

func RandInt(n int) []int {
	rand.Seed(time.Now().Unix())
	var list []int
	for i := 0; i < n; i++ {
		list = append(list, rand.Intn(n))
	}
	return list
}
