package constant_test

import "testing"

// 一组常量中，如果某个常量没有初始值，默认和上一行一致
// iota，特殊常量，可以认为是一个可以被编译器修改的常量
// 每当定义一个const，iota的初始值为0,每当定义一个常量，就会自动累加1

const (
	one   = iota //iota初始值为0
	two   = iota //自增1
	three = iota //再自增1
)

func TestConstantTry0(t *testing.T) {
	t.Log("one 的值为：", one)
	t.Log("two 的值为：", two)
	t.Log("three 的值为：", three)
}

const (
	Monday = iota
	Tuesday
	Wednesday
)

const (
	Readable   = 1 << iota // 表示（0001）
	Writable               // 表示（0010）
	Executable             // 表示（0100）
)

func TestConstantTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7                                // 0111
	t.Log(Readable, Writable, Executable) // 1 2 4
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
