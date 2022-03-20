package operator_test

import "testing"

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b) // false
	// t.Log(a == c)
	t.Log(a == d) // true
}

func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readable
	t.Log(a) // 6
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}

// z = x &^ y
// 如果运算符右侧数值的第 i 位为 1，那么计算结果中的第 i 位为 0；
// 如果运算符右侧数值的第 i 位为 0，那么计算结果中的第 i 位为运算符左侧数值的第 i 位的值。
