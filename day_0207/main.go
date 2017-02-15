// test2 project main.go
package main

import (
	"fmt"
)

// 1

func main() {
	var b bool
	var n int //定义变量n
	var (
		aa  int = 3
		str string
	)

	var i1, i2, i3 int = 1, 2, 3
	var strName = "张三"
	var strSex = "男"
	//fmt.Println("hello")

	fmt.Println("n=", n)
	fmt.Println("b=", b)
	fmt.Println("aa=", aa)
	fmt.Println("str=", str)
	fmt.Println("i1=", i1)
	fmt.Println("i2=", i2)
	fmt.Println("i3=", i3)
	fmt.Println("strName=", strName)
	fmt.Println("strSex=", strSex)

}

/*
数组的引用
*/
/*
func main() {
	arr1 := [2]int{}
	arr2 := [2]int{1, 2}
	arr3 := [2]int{0: 1, 1: 2}
	arr4 := [...]int{1, 2}
	arr5 := [...]int{3: 9} //声明一个4个元素的数组，
	//名称为arr5，并赋值为[0,0,0,9]

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr5)

}
*/
/**
切片  Slice
和数组的声明语法类似，但是不像数组那样要指定元素的个数
有弹性，效率高
[]T
*/
