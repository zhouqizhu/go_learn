package interface_test

import "testing"

// Go 接口是一组方法的集合，可以理解为抽象的类型。
// 定义Programmer接口
type Programmer interface {
	WriteHelloWorld() string
}

// 定义GoProgrammer结构体
type GoProgrammer struct {
}

// 方法是指和结构体绑定的，即func (u *User) method() {}这样带有接收器的。
func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World\")"
}
// Duck Type 式接口实现
func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 和其他编程语言区别：
// 1.接⼝口为⾮非⼊入侵性，实现不不依赖于借⼝口定义
// 2. 所以接⼝口的定义可以包含在接⼝口使⽤用者包内
