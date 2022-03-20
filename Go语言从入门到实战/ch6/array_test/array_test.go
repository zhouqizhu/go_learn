package array_test

import "testing"

func TestArrayInit(t *testing.T) {
	var arr [3]int
	arr1 := [4]int{1, 2, 3, 4}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr1[1] = 5
	t.Log(arr[1], arr[2])
	t.Log(arr1, arr3)
}

// 数组遍历
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	// idx可以用_占位
	for idx, e := range arr3 {
		t.Log(idx, e)
	}
}

//数组截取
func TestArraySection(t *testing.T) {
	// a[开始索引(包含), 结束索引(不包含)]
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr3_sec := arr3[:3]
	t.Log(arr3_sec)
}
