// day_0215_CommandPro project main.go
package commandPro

//例子网址：
//https://hacpai.com/article/1483778326912
//接口

import (
	//"day_0215_CommandPro/collection"
	"fmt"
)

type MyType struct{ i int }
type MyInt int

var pm = new(MyType) //定义变量pm结构体
var n = pm.Get()

//接口
type MyInterface interface {
	Get() int
	Set(i int)
}

//实现接口方法
//func (p MyInt) Get() int {
//	return int(p) //强制转换
//}
func (p *MyType) Get() int {
	return p.i
}
func (p *MyType) Set(i int) {
	p.i = i
}

func GetAndSet(x MyInterface) {}

func f1() {
	var p MyType
	GetAndSet(&p)
	p.Set(45)
	fmt.Println("i:", p.Get())
}

type MySubType struct {
	MyType
	j int
}

//匿名字段的使用  这里重写Get方法
func (p *MySubType) Get() int {
	p.j++
	return p.MyType.Get()
}

//这里Set方法则被继承
func f2() {
	var p MySubType
	GetAndSet(&p)
	p.Set(67)
	fmt.Println("i Set:", p.Get())
	fmt.Println("p.j:", p.j)
}

type Printer interface {
	Print()
}

//一个接口类型的变量可能使用类型断言转换为不同的接口类型
func f4(x MyInterface) {
	x.(Printer).Print() //!看不懂
}

func main() {
	fmt.Println("fmtee3")
	//f1()
	f2()
}
