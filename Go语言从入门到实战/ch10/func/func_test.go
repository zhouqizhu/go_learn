package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 可以有多个返回值
func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(90)
}

// 函数可以作为参数和返回值
// 函数可以作为变量量的值
// inner 的类型是func(opt int) int
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

// 所有参数都是值传递：slice，map，channel 会有传引⽤用的错觉
func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10)) // 10
}

// 可变参数
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))    // 10
	t.Log(Sum(1, 2, 3, 4, 5)) // 15
}

// defer(延迟执行语句)
func Clear() {
	fmt.Println("Clear resources.")
}
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Println("start.")
	panic("err") // 程序终止运行
	fmt.Println("end")
}
