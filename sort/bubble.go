package sort

// BubbleWithInt 冒泡排序 最慢的排序算法
func BubbleWithInt(list []int32) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]

			}
		}
	}
}

// BubbleWithInt64 最慢的排序算法
func BubbleWithInt64(list []int64) {
	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
