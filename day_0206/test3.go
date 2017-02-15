// test3 project test3.go
package main

import "fmt"

//import "reflect"

func main() {
	p := [...]int{2, 3, 5, 7, 11, 13} //定义一个数组
	s1 := p[1:3]                      //定义切片

	fmt.Println(s1)
}
