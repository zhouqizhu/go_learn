package main //包，声明代码所在的模块（包）

//引入代码依赖
import (
	"fmt"
	"os"
)

//功能实现
// func main() {
// 	fmt.Println("Hello World")
// 	// return 1  报错，main方法不能有返回值
// 	os.Exit(-1)
// }

func main() {
	if len(os.Args) > 1 {
		fmt.Println("hello wordld", os.Args[1])
	}
}

//注意：
//应用程序入口必须是main包：package main
//必须是main方法：func main()
//文件名不一定是main.go

//与其它语言差异：
// main函数不支持传入参数
// 在程序中直接通过os.Args获取命令行参数
//Go中main函数不支持任何返回值
//通过os.Exit来返回状态
