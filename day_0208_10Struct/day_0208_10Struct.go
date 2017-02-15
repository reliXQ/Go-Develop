// day_0208_10Struct project day_0208_10Struct.go
package main // day_0208_10Struct

import "fmt"

/*
	go中没有继承的概念
	struct与C 中的很像
	go 没有class
	定义：  type<Name> struct{}
	支持指向自身指针类型成员
	支持匿名结构，可用作成员或定义成员变量
    嵌入结构作为匿名字段看起来像继承，但不是继承

	//匿名结构可以嵌套在其他格式当中
*/

//01
/*
type person struct {
	Name string
	Age  int
}

func main() {

	//		a := person{}
	//		a.Name = "joe"
	//		a.Age = 19
	//		fmt.Println(a)

	//字面值的初始化
	a := person{
		Name: "joe",
		Age:  19,
	}
	fmt.Println(a)
	A(a) //说明这里是值得拷贝  跟c语言一样  如果想改变其值，只需要传入其地址
	fmt.Println(a)
}

func A(per person) {
	per.Name = "John"
	per.Age = 20
	fmt.Println("A:", per.Name, per.Age)
}

//推荐结构体这样定义  推荐结构初始化的时候加上&符号
//a:=&person{  //这里加了&符号
//	Name:"joe",
//	Age:19,
//}
//  a.Name = "OK"  //一样可以用

*/

//02
//匿名结构  匿名结构嵌套在其他格式当中
/*
type person struct {
	Name    string
	Age     int
	Contact struct {
		Phone string
		City  string
	}
}

func main() {
	a := &struct { //匿名结构
		Name string
		Age  int
	}{
		Name: "joe", //! ","
		Age:  19,
	}
	fmt.Println(a)
	//匿名结构嵌套在其他格式当中
	b := person{
		Name: "UUU", Age: 23,
	}
	b.Contact.Phone = "1233334"  //结构体中的匿名结构只能通过这种方法赋值
	b.Contact.City = "beijing"
	fmt.Println(b)
}
*/

//03
//匿名字段
/*
type person struct {
	string
	int
}

func main() {
	a := person{"joe", 19} //!  {}写成了（）  这里的顺序不能错
	var b person
	b = a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(b == a) //比较的是地址
	//只能相同类型之间进行比较  如果地址相同，则比较内容是否相同

}
*/

type human struct {
	Sex    int
	Height int
}

type teacher struct {
	human //组合  go中不存在泛型
	Name  string
	Age   int
}
type student struct {
	human //组合
	Name  string
	Age   int
}

func main() {
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "iii", Age: 29, human: human{Sex: 1}}
	//嵌入结构
	a.Age = 13
	a.Name = "joe2"
	a.Sex = 99 //
	a.human.Sex = 2
	fmt.Println(a, b)

}
