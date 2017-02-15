// day_0207_type project day_0207_type.go
package main

//import "fmt"

/*
type (
	byte     int8
	rune     int32
	ByteSize int64
)
*/
/*
type (
	newa int
)

var () //全局变量

func main() {
	var a float32 = 100.12
	fmt.Println(a)
	b := int(a)
	fmt.Println(b)
}
*/

//二
/*
import "math"
import "fmt"

func main() {
	fmt.Println(math.MaxInt32)
}
*/

//三
/*
import "fmt"

//这里声明变量为全局变量
var a int

//a = 123   //变量的赋值

func main() {
	//a = 23
	//变量声明的同时赋值
	var b int = 321 //Wrong!  var b int 321
	//省略变量类型 系统判断
	var c = 32.5
	//变量声明和赋值的最简写法
	d := 12
	a = 34
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
*/
// end

//四
/*
import "fmt"

func main() {
	//多个变量的声明
	var a, b, c, d int
	//多个变量的赋值
	a, b, c, d = 1, 2, 3, 4

	//多个变量声明的同时赋值
	var e, f, g int = 1, 3, 5
	//省略变量类型有系统推断
	var h, i, j = 5, 6, 7
	//多个变量类型的最简写法
	//b, c, d := 23, 45, 1111  //Wrong! no new variables on left side of :=
	//":=" 这个是定义变量用的  不能赋值的时候用
	b, c, d = 23, 45, 1111

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

}
*/

//5

import "fmt"

func main() {
	a, b, _, d := 1, 2, 3, 4 //"_"此方法一般用在 函数多参数返回的时候
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
	fmt.Println("-----")

	//强制转换
	var e float32 = 23.45
	fmt.Println(e)
	//b := bool(e)  //Wrong!  float32不能强制转换为bool
	f := int(e)
	fmt.Println(f) //float32 转 int

	var g int = 65
	h := string(g)
	fmt.Println(h) //打印的是65的ascii值
	//下一节 告诉你怎么打印String 65
}
