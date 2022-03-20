package slice_test

import "testing"

// 切片定义
func TestSliceInit(t *testing.T) {
	var s0 []int
	// len: 元素的个数
	// cap: 内部数组的容量
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)      // 0 0
	t.Log(s0)               // [1]
	t.Log(len(s0), cap(s0)) // 1 1

	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1)) // 4 4

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2))    // 3 5
	t.Log(s2[0], s2[1], s2[2]) // 0 0 0
	s2 = append(s2, 1, 2, 3)
	t.Log(s2[0], s2[1], s2[2], s2[3], s2[4], s2[5]) // 0 0 0 1 2 3
	t.Log(len(s2), cap(s2))                         // 3 10
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // 1 1 2 2 3 3 4 4 5 8 6 8 7 8 8 8 9 16 10 16
	}
}

// 切片共享存储结构
func TestSliceShareMemory(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2)) // [Apr May Jun] 3 9
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer)) // [Jun Jul Aug] 3 7
	summer[0] = "Unknow"
	t.Log(Q2)   // [Apr May Unknow]
	t.Log(year) // [Jan Feb Mar Apr May Unknow Jul Aug Sep Oct Nov Dec]
}

// 切片只能和空nil比较，不能和切片比较
// func TestSliceComparing(t *testing.T) {
// 	a := []int{1, 2, 3, 4}
// 	b := []int{1, 2, 3, 4}
// 	if a == b {
// 		t.Log("equal")
// 	}
// }

// 数组和切片区别
// 1.容量是否可以伸缩
// 	数组不可伸缩
// 	切片可以伸缩
// 2.是否可以进行比较
// 	数组 相同长度元素，相同元素的数组可以比较
