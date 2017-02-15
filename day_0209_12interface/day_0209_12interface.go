// day_0209_12interface project day_0209_12interface.go
package main // day_0209_12interface

import "fmt"

/*
	接口interface
	接口是一个或多个方法签名的集合
	Structural Typing
	只要某个类拥有了该接口的所有方法签名，既算是实现该接口，无需显示声明实现了哪个接口，这称为Structural Typing
	接口只有方法声明，没有实现，没有数据字段
*/

//01
//接口
//这是个usb接口  如果想实现此接口  则必须实现此接口中的方法就可以了
/*
type USB interface {
	Name() string //
	Connecter     //  接口嵌套
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct { //PhoneConnecter  首先定义个结构体他实现了接口 USB接口
	name string
}

func (pc PhoneConnecter) Name() string { //pc:PhoneConnector实现了USB的Name（）接口
	return pc.name //返回的是自身的名字  并不是接口中的名字//pc.Name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	//方法1
	//	var a USB //定义一个USB接口变量
	//	a = PhoneConnecter{"TelePhone"}    //!W   cannot convert "TelePhone" (type string) to type PhoneConnecter  "()"换成“{}”:对结构体中变量的赋值方法掌握不准确
	//方法2
	a := PhoneConnecter{"TelePhone"}
	a.Connect()
	Disconnect(a)
}

//这就是简单的接口断言
//方法1
//func Disconnect(usb USB) {
//	//方法1
//	//	fmt.Println("Disconnected") //问题： 这里因为USB接口中没有.name 参数  我们现在取不到
//	//方法2
//	if pc, ok := usb.(PhoneConnecter); ok { //这里我们判断ok 是否成立
//		fmt.Println("Disconnected:", pc.name)
//		return
//	}
//	fmt.Println("Unknow device")
//}

//方法2  使用空接口改造  这种使用的比较多
func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknow device")
	}
}

//分析：go中没有继承
//通过接口 默认承认存在上帝之源  ==等同于  继承
*/

//02
//重写上面的接口事例
/*
//接口
type USB interface {
	Name() string //!W Name string   ???
	Connecter     //嵌入式接口应用     //Connect()     //  接口中的方法都要加（）
}
type Connecter interface {
	Connect()
}
type PhoneConnecter struct {
	name string
}

//实现接口
func (pc PhoneConnecter) Name() string { //!少了string
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

//断开连接函数
//1
//func Disconnect(usb USB) {
//	fmt.Println("Disconnect", usb)
	if pc,ok :=usb.(PhoneConnecter); ok{  //pc 是转换后的结果  第二个：ok 是转换是否成功  如果ok是false的话
	//则pc一般是0值或是空
		fmt.Println("Disconected:",pc.name)
		return
	}
	fmt.Println("Unknow device.")
//}
//2
func Disconnect(usb USB) {
	switch v := usb.(type) {  //大量的应用  推荐！
 	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknow device")
	}
}
func main() {
	var a USB //定义一个a实现USB接口
	a = PhoneConnecter{"Teleconnect"}
	a.Connect()
	Disconnect(a)
}
*/
/*
问题1：
 cannot use PhoneConnecter literal (type PhoneConnecter) as type USB in assignment:
	PhoneConnecter does not implement USB (missing CCon method)
答： 这是因为没有完成接口的所有初始化导致的  完成所有接口即可
*/

//03
type USB interface {
	Name() string
	Connecter
}
type Connecter interface {
	Connect()
}

//结构体
type PhoneConnecter struct {
	name string
}

//实现接口
func (pc PhoneConnecter) Name() string {
	return pc.name
}
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connected:", pc.name)
}

//断开连接函数
func Disconnect(usb USB) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknow device")
	}
}
func main() {
	pc := PhoneConnecter{"Teleconnect"}
	var a Connecter
	a = Connecter(pc) //把PhoneConnecter强制转换为Connecter的的  不太懂这样用的方法有什么用
	//相当于a继承了一个PhoneConnecter的子类接口  它里面只有Connecter这一个方法
	a.Connect()
	//Disconnect(a)

	fmt.Println("----")
	var b interface{}
	fmt.Println(b == nil)

	var p *int = nil
	b = p //b指向的类型不是nil  是个指针  所以不是true
	fmt.Println(b == nil)

}

//注意：
//接口的调用不会做 receive的自动转换
//接口  他是个指针类型 你就必须传递一个指针给他
