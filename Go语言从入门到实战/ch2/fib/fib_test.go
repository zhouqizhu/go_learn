package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	// var a int = 1
	// var b int = 1

	// var (
	// 	a int = 1
	// 	b int = 1
	// )
	// := 是声明并赋值，并且系统自动推断类型，不需要var关键字
	a := 1
	b := 1
	// fmt.Print(a)
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

// 交换两个变量的值
func TestExchange(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	a, b = b, a
	t.Log(a, b)
}
