package condition_test

import "testing"

func TestIfMultiSec(t *testing.T) {
	// ⽀持变量赋值
	// if  var declaration;  condition {
	// code to be executed if condition is true
	// }
	if a := 1 == 1; a {
		t.Log("1==1")
	}

	// if v, err := someFun(); err == nil {
	// 	t.Log("1==1")
	// } else {
	// 	t.Log("1==2")
	// }
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		// 1. 条件表达式不限制为常量或者整数；
		// 2. 单个 case 中，可以出现多个结果选项, 使⽤逗号分隔；
		// 3. 与 C 语⾔等规则相反，Go 语⾔不需要⽤break来明确退出⼀个 case；
		// 4. 可以不设定 switch 之后的条件表达式，在此种情况下，整个 switch 结构与多个 if…else… 的逻辑作⽤等同
		// switch i {
		// case 0, 2:
		// 	t.Log("Even")
		// case 1, 3:
		// 	t.Log("Odd")
		// default:
		// 	t.Log("it is not 0-3")
		// }

		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unknow")
		}
	}
}
