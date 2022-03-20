package map_ext

import "testing"

// Map 的 value 可以是⼀个方法
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2)) // 2 4 8
}

// Set元素唯一
func TestMapForSet(t *testing.T) {
	// Go 的内置集合中没有 Set 实现， 可以 map[type]bool
	// mySet1 := map[int]bool{1: true, 2: true, 2, true}
	// t.Log(mySet1)报错
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n) //  3 is not existing
	}
	mySet[3] = true
	t.Log(len(mySet)) // 2
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n) // 1 is not existing
	}
}
