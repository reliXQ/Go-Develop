// day_0207_05Array project day_0207_05Array.go
package main //day_0207_05Array

import "fmt"

// 01
//数组
//var <varName>[n]<type> , n>=0
//指向数组的指针和指针数组

func main() {
	var array [4]int
	var arr [3]int
	//赋值
	a := [2]int{1} //赋值第一个元素为“1”

	fmt.Println(array)
	fmt.Println(arr)
	fmt.Println(a)

	//特殊赋值  索引
	b := [20]int{19: 1} //第20个元素为1
	fmt.Println(b)
	fmt.Println("---")
	c := [...]int{2, 3, 4}
	fmt.Println(c)

	fmt.Println("---")
	d := [...]int{1: 2, 5: 4, 20: 12}
	fmt.Println(d)

	fmt.Println("--1-------------")
	e := [...]int{17: 1}
	var f *[18]int = &e //指向数组的指针
	fmt.Println(f)
	fmt.Println("--2--------------")
	x, y := 1, 2
	g := [...]*int{&x, &y} //指针数组
	fmt.Println(g)
	fmt.Println("-----------------")

	o := [10]int{} //int后面一定要加{}
	o[1] = 2
	p := new([10]int) //指向数组的指针
	p[1] = 2
	fmt.Println(p)

	fmt.Println("---")
	q := [2][3]int{ //2行3列
		{1, 1, 1},
		{2, 2, 2}, //记得要加  “，”
	}
	fmt.Println(q)
}

//go语言版的冒泡排序
/*
func main() {
	a := [...]int{4, 2, 6, 1, 7}
	fmt.Println(a)
	count := len(a)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
*/
