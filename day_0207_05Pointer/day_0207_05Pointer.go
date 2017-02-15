// day_0207_05Pointer project day_0207_05Pointer.go
//package day_0207_05Pointer
package main

import "fmt"

/*
//01 指针
//不支持指针的运算，而是直接采用类中的“.”来操作指针目标对象的成员
//&取地址   *取对象  默认值是：nil 不是NULL  没有指向任何地址
//go语言当中 ++与--是作为语句而并不是作为表达式

func main() {
	a := 1
	var p *int = &a //定义一个p（int型的指针）指向a的地址
	fmt.Println(p)  //打印的是地址值
	fmt.Println(*p)

//		b := 2
//		//var q *int = a  //这样写是错的  cannot use a (type int) as type *int in assignment
//		fmt.Println(q) //*int 指针 不能指向一个变量q  只能指向变量的地址
//		fmt.Println(&q)


}
*/
/*
func main() {
	if a := 1; a > 2 { // syntax error: missing { after if clause
		//缺少一个“;”  //应该是：a:1;a>2
		//这里的a 是一个局部变量
		fmt.Println("true")
	}
	fmt.Println(2)
}
*/

//02 只有for 一个关键字 没有while
//有三种形式  其中有一种是可以代替while功能的
/*
func main() {
	a := 1
	//方式1
	//	for {
	//		a++
	//		if a > 3 {
	//			break
	//		}
	//		fmt.Println(a)
	//	}
	//	fmt.Println("over")

	//方式2
	for a <= 3 {
		a++
		fmt.Println(a)
	}
	fmt.Println("over2")

	//方式3
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("over3")

	//注意： 不要在 for 的函数表达式中设置函数  不然每次执行循环都回去计算一次
	//e.g:  for i := 0; i < len(a); i++ {}   //！！不要这样写

}
*/

//03
//switch 用法
//不需要写break   一旦条件符合自动终止
//如需继续执行下一个case ，需要 “fallthrough”语句
//支持一个初始化表达式（可以是并行方式），右侧需要跟分号
/*
func main() {

	//a := 0
	//	switch a {
	//	case 0:
	//		fmt.Println("a=0")
	//	case 1:
	//		fmt.Println("a=1")
	//	}
	//		default:
	//			fmt.Println("None")
	//	fmt.Println(a)

	//	switch a {
	//	case 0:
	//		fmt.Println("a=0")
	//		fallthrough  //作用不好  //现象：a=0  a=1
	//	case 1:
	//		fmt.Println("a=1")

	switch a := 1; { //这里的a:=1;是初始化的一个变量  判断表达式必须要写在下面的case中
	case a <= 0: //！case 0:
		fmt.Println("cast = 0")
		fallthrough //作用不太好
	case a <= 1:
		fmt.Println("case = 1")
		fallthrough
	default:
		fmt.Println("None")
	}

}
*/

//break  通过标签可以跳出多层循环
