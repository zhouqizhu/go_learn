package string_test

import "testing"

func TestString(t *testing.T) {
	var s string
	// 初始化为默认零值""
	t.Log(s)
	s = "hello"
	t.Log(len(s)) // 5
	//s[1] = '3' //string是不可变的byte slice
	s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s)) // 4

	s = "中"
	t.Log(len(s)) // 3
	// rune 等同于int32,常用来处理unicode或utf-8字符
	c := []rune(s)
	t.Log(len(c)) // 1

	t.Logf("中 unicode %x", c[0]) // 中 unicode 4e2d
	t.Logf("中 UTF8 %x", s)       // 中 UTF8 e4b8ad
}
func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c) // 中 4e2d // 华 534e // 人 4eba // 民 6c11 // 共 5171 // 和 548c // 国 56fd
	}
}

// 1. string 是数据类型，不是引⽤或指针类型
// 2. string 是只读的 byte slice，len 函数可以它所包含的 byte 数
// 3. string 的 byte 数组可以存放任何数据
