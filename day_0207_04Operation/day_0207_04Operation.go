// day_0207_04Operation project day_0207_04Operation.go
//package day_0207_04Operation
package main

import "fmt"

//类名的空值  无定义的值
//类名的别名	新的名称

//申请全局变量可以用
//var （
//	一次性声明很多变量
//	aa int = 3
//	b string
//）
//变量内部声明则不可以用这种方法//!但我测试了是可以的

//全局和局部声明都可以用并行的方式
/*

import "strconv"

func main() {
	var a int = 65
	b := strconv.Itoa(a) //把int65 转成字符串
	fmt.Println(b)
	var c int
	c, _ = strconv.Atoi(b) //!  这里不能只写c  //报错： multiple-value strconv.Atoi() in single-value context  具有多个值
	fmt.Println(c)
}
*/

//实例2
/*
//常量的定义
//和变量定义的格式基本相同

const a int = 1
const b = 'A'
const (
	text   = "123"
	length = len(text)
	num    = b * 20
)

//同时定义多个变量
const i,j,k =1,"2","3"
const(
	text2,length2,num2 = "345",len(text2),k*10
)
//如果不给常量赋值，则他直接使用上一行的表达式的值
*/

//实例3
/*
const (
	a = "A"
	b
	c = iota  //iota 和这个常量的顺序有关
	d
	MAX_CONT  //常量大写（最好）
	//首字母如果是大写的会被外部调用到，如果你不想把这个常量
	//可以在首字母前再加“_”
	//e.g:  _MAX_CONT  或  cMAX_CONT
)
const (
	e = iota //iota 是常量的计数器，没遇到一次const
	f        //它就会重置为0，每定义一个常量就自动加1（从0开始）
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("----")
	fmt.Println(e)
	fmt.Println(f)
}
*/

//枚举
/*
const (
	Monday = iota
	Tuesday
	Wednewday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednewday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
	fmt.Println(Sunday)
}
*/

//运算符
//优先级  从左到右
/*
func main() {
	fmt.Println(1 << 10)       //JG: 1024  //二进制左移10位 按位移动
	fmt.Println(1 << 10 << 10) //JG：1048576

	a := 0
	if a > 0 && 10/a > 1 { //多了一个判断  “且”
		fmt.Println("ok")
	}

}
*/

//作业： 结合常量的iota与<<运算符实现计算机存储单位的枚举
const (
	B float64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	//	fmt.Printf("B:[%lf]", B)  //！这样写有问题
	//	fmt.Printf("KB:%lf", KB)  //B:[%!l(float64=1)f]KB:%!l(float64=1024)f
}
