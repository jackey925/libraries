package sort

import (
	"github.com/jackey925/libraries/random"
	"sort"
	"testing"
)

// BubbleWithInt 冒泡排序
func TestBubbleWithInt(t *testing.T) {
	s := random.RandInt32(30)
	t.Log(s)
	BubbleWithInt(s)
	t.Log(s)

}

func BenchmarkBubbleWithInt(b *testing.B) {
	s := random.RandInt32(10000)
	for i := 0; i < b.N; i++ {
		BubbleWithInt(s)
	}
}

func BenchmarkSort(b *testing.B) {
	s := random.RandInt(10000)
	for i := 0; i < b.N; i++ {
		sort.Ints(s)
	}
}
