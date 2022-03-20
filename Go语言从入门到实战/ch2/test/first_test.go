// 编写测试文件
// 源码文件以_test结尾：xxx_test.go
// 测试方法名以Test开头：func TestXXX(t *testing.T){...}
package try_test

import "testing"

// *是指针运算符 , 可以表示一个变量是指针类型 , 也可以表示一个指针变量所指向的存储单元 , 也就是这个地址所存储的值
func TestFirstTry(t *testing.T) {
	t.Log("My first try!")
}

// 与其他主要语言差异：
// 赋值可以进行自动类型推断
// 在一个赋值语句中可以对多个变量进行同时赋值
