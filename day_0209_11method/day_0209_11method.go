// day_0209_11method project day_0209_11method.go
package main // day_0209_11method

import "fmt"

/*
复习Struct
两种想通的结构才能赋值
嵌入结构
问题：如果内层结构和外层结构中都有同名字段怎么办？

method

*/

/*
//01
//上一课的Struct 作业如下：  如果A:中有Name 嵌套的B中也有Name怎么去赋值？？？
type A struct { //a.Name  如果A结构体中没有所需要的.Name 则会往下一级别找.Name字段
	B
	Name string
}
type B struct {
	Name string
}

func main() {
	//	a := A{Name: "A", B: {Name: "B"}}  // !W missing type in composite literal
	a := A{Name: "A", B: B{Name: "B"}}
	fmt.Println(a.Name, a.B.Name)

}
*/

//02
//方法的定义方法
/*
type A struct {
	Name string
}
type B struct {
	Name string
}

func main() {
	a := A{}  //！A() 错误: missing argument to conversion to A: A()
	a.Print() //写一个Print方法
}

//method  编译器根据接收者的类型来判断是那种方法
//go 中 没有方法重载 的概念
func (a A) Print() { //a:接收者  A:类型  Print：方法名称  ():参数列表    //这就是一个方法
	fmt.Println("A")
}
*/

/*
//03
// *A  在method方法中修改类型的值  需要加*  ：这样就可以修改结构体的数据了
type A struct {
	Name string
}
type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

}

//如果想在method中修改struct中的值  必须传进去地址
func (a *A) Print() {
	a.Name = "AA" //这个修改 Struct中a.Name成功 //这里a 是系统自动识别的 不必*a
	fmt.Println("A")
}

func (b B) Print() { //这里b只是得到了一个副本
	b.Name = "BB"
	fmt.Println("B")
}
*/

//04
/*
//首字母大写是导出字段，小写是私有字段
type TZ int // type 把一个类型的名称 转化成另外一个类型而保持数据本质不变
//type 的这种绑定 可以为任意类型的 数据绑定方法
//有点面相对象的思想
func main() {
	//a := A{}
	//a := TZ //!W  not a expression
	var a TZ
	a.Print()
}

//方法绑定
func (a *TZ) Print() { //这里的a 是个变量  类型：TZ（int）   也就是说一个变量可以有函数方法
	fmt.Println("A TZ")
}
*/

//05
//作业： 声明一个底层类型为int的类型，并实现调用某个方法就递增100
//如：a:=0 调用a.Increase()后，a从0变成了100
//自己思考：
/*
type TZ int

func main() {
	var a TZ
	a.Increase()
	fmt.Println(a)

}

func (a *TZ) Increase() int {
	a = (*TZ)100;
	return a
}
//没完成  报错  找不到问题的原因
//go:123: syntax error: unexpected literal 100 at end of statement
*/
//老师答案：
//写法1
/*
type TZ int

func main() {
	var a TZ
	a.Increase()
	fmt.Println(a)

}

func (a *TZ) Increase() TZ {
	//	*a += (TZ)100  //go:141: syntax error: unexpected literal 100 at end of statement
	*a = TZ(100) //强制类型转换的写法是这样的
	return *a
}
*/

//方法2
type TZ int

func (a *TZ) Increase(num int) {
	*a += TZ(num)
}
func main() {
	var a TZ
	a.Increase(100)
	fmt.Println(a)
}
