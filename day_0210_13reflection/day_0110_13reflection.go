// day_0210_13reflection project day_0110_13reflection.go
package main //day_0210_13reflection

import "fmt"

/*
	反射 ：灵活
	使interface更具有灵活性
	TypeOf和ValueOf函数  获取目标对象的信息
	反射会将匿名字段作为独立字段
	如果想用反射修改对象的状态，前提：interface.data 是settable  既
	pointer-interface
	通过反射动态的调用方法
*/

//01
//reflection
//TypeOf  ValueOf
/*
import "reflect"

type User struct {
	Id   int
	Name string
	Age  int
}

//func (u User) Hello() { //这个结构体有个方法  实现它
//	fmt.Println("Hello nihao")
//}

func main() {
	u := User{1, "OK", 12} //初始化结构体
	Info(u)                //调用方法
}

//方法： 传进去一个空接口(任意类型)  输出我们传进去"结构"的具体信息
func Info(o interface{}) {
	t := reflect.TypeOf(o) //类型
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o) //包含的字段
	fmt.Println("Fields:")

	if k := t.Kind(); k != reflect.Struct { //k:=t.Kind()  取出类型
		fmt.Println("！W  不是XX类型的")
		return
	}
	//取得所有字段信息以及类型信息
	for i := 0; i < t.NumField(); i++ { //NumField 数字领域
		f := t.Field(i)                                  //取出字段
		val := v.Field(i).Interface()                    //取出字段所对应得值
		fmt.Printf("%8s:%v = %v\n", f.Name, f.Type, val) //名称  类型  字段的值
	}

	//取得方法的信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("N:%6s T:%v F:%s I:%d  P:%s ", m.Name, m.Type, m.Func, m.Index, m.PkgPath)
	}
}

//对某一些结构进行反射
*/

//02
//下面  反射结构中的匿名  或嵌入字段
/*
import "reflect"

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User  //匿名字段会当成一个独立的字段来处理
	title string
}

func main() {
	m := Manager{User: User{1, "Mana", 11}, title: "M"}
	//u := User{1, "OK", 12} //初始化结构体

	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0)) //打印结构体的信息
	//结果：reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x4a2800), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	fmt.Println("----")
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Println("----")
	//fmt.Printf("%#v\n", t.Field(2)) //!W  panic: reflect: Field index out of bounds  超出字段 不存在这个字段
	//怎么取匿名字段中的数据
	fmt.Println("---id-")
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0})) //传进去一个  []int{0,0} :int型的slice来取匿名字段中的数据
	fmt.Println("---Name-")
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
}
*/

//03
//反射 某个结构的信息 某个类型的信息
//下面介绍 怎么通过反射修改它的内容
/*
import "reflect"

func main() {
	x := 123
	v := reflect.ValueOf(&x)  //取出它（x）的值  reflect。point
	v.Elem().SetInt(999) //Elem取得对应的值

	fmt.Println(x)
}
*/

//04
//通过接口中的对象，对它的值进行修改
/*
import "reflect"

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	u := User{1, "ok", 12}
	Set(&u) //传一个指针给这个函数
	fmt.Println(u)
}

//需要传进来 指针类型的 值 才能修改
func Set(o interface{}) {
	v := reflect.ValueOf(o) //取出值      //这里的v是 reflect.value类型
	//判断是否是 point interface 的语句是 ：v.Kind() == reflect.Ptr
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //!W  "==" 写成了“=”  ：syntax error: missing { after if clause
		//并且去掉了“k=v.Kind()”中的k=
		fmt.Print("xxx  不能修改")
		return
	} else {
		v = v.Elem() //重新赋值   //这里的v也是 reflect.value类型
	}

	//注释：v.FieldByName("Name")  //取出字段 ：通过（名称）取出要返回的值
	//注释：f.Kind()==reflect.String 判断是不是reflect.String类型
	//方法1：
	//	if f:=v.FieldByName("Name");f.Kind()==reflect.String{
	//		f.SetString("BABY")
	//	}
	//方法2：
	f := v.FieldByName("Name") //赋值
	if !f.IsValid() {          //判断是否真的找到了这个字段  如果此字段不存在  打印“Not Exist” 并返回
		fmt.Println("Not Exist")
		return
	}

	if f.Kind() == reflect.String {
		f.SetString("BABY")
	}
}
*/

//05
//如何通过反射进行方法的调用
//动态调用这个方法
/*
import "reflect"

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello(name string) {
	fmt.Println("Hello", name, "my name is ", u.Name)
}
func main() {
	u := User{1, "ok", 12}
	//u.Hello("joe")

	v := reflect.ValueOf(u) //创建反射对象
	mv := v.MethodByName("Hello")
	//建立一个切片
	args := []reflect.Value{reflect.ValueOf("joe")}
	//reflect.string   reflect.int  比他们更高一级的：reflect.Value

	//调用方法
	mv.Call(args)
}
*/
