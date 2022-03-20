package type_test

import "testing"

// 类型定义,通过 type 关键字的定义，MyInt会形成一种新的类型，MyInt本身依然具备 int 类型的特性。
type MyInt int64

// 类型别名
// type MyInt = int64

func TestImplicit(t *testing.T) {
	// 1. Go 语⾔不允许隐式类型转换
	// 2. 别名和原有类型也不能进行隐式类型转换
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	// & 是取地址符号 , 即取得某个变量的地址
	aPtr := &a
	// ×  aPtr = aPtr + 1,不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // 格式化输出其类型
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") // **
	// string 是值类型，其默认的初始化值为空字符串，而不是 nil
	t.Log(len(s)) // 0
}
