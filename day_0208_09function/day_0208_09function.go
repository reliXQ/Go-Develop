// day_0208_09function project day_0208_09function.go
package main // day_0208_09function

import "fmt"

/*
	go 函数不支持嵌套，重载，默认参数
	但支持
		无需声明原型，不定长度变参，多返回值，命名返回值参数，匿名函数，闭包
*/

//01
//函数定义
/*
func main() {

}

func A(a int, b string) (int, string, int) { //第一个（）是参数  第二个（）是返回值类型

}

func B(a int, b int) int { //返回值只有一个  int

}

// 不定长变参
func C(a, b, ...int) (a, b, c int) { //...不定长变参
	return a,b,c
}
*/

//02
//案例
/*
func main() {
	s1 := []int{1, 2, 3, 4, 5} //切片
	a := 0
	funcA(s1)
	fmt.Println(s1)
	fmt.Println("-----")
	funcB(a)
	fmt.Println(a)
	fmt.Println("-----")
	funcC(&a)
	fmt.Println(a)
	fmt.Println("-----")
}

func funcA(s []int) {
	s[0] = 5
	s[1] = 6
	s[2] = 7
	s[3] = 8
	fmt.Println(s)
}
//结论：切片穿进去的修改 外部的实际数组值  相当于传的是形参
func funcB(a int) {
	a = 9
	fmt.Println(a)
}
//结论：变量传进去 相当于是赋值  外部值不变
func funcC(a *int) {
	*a = 8
	fmt.Println("传递进来指针")
	fmt.Println(*a)
}
//结论：此方法传进去变量的指针，相当于变量的形参

//结论： 函数传递进去值不改变，传递进去切片slice  改变被传递的slice
//原因： 被传递的slice是传递的形参（地址的拷贝）  而传递的值是以数值传递进去的值得拷贝
//如需想改变值 则只需要传递进去值得地址即可
//go语言中 “一切皆类型”
*/

//03
//匿名函数 & 闭包
//go语言中一切皆函数
/*
func main() {
	a := func() {  //这是匿名函数
		fmt.Println("Func A")
	}
	a() //调用A
}
*/
/*
func main() {
	f := closure(10)  //闭包
	fmt.Println(f(1)) //！W   cannot use f(1) (type int) as type string in argument to fmt.Printf  //错误原因：误写为：Printf
	fmt.Println(f(2))

}

func closure(x int) func(int) int {
	fmt.Printf("A: %p\n", &x) //说明传进来的是地址    %p指针
	return func(y int) int {  //传进来一个y  返回int类型数据
		fmt.Printf("B: %p\n", &x)
		return x + y //x=10  y每次不同是1,2   多次使用了x=10
	}
}
*/
/* 结果
A: 0xc0420381d0
B: 0xc0420381d0
11
B: 0xc0420381d0
12
*/

//04
//defer
//析构函数
//在函数执行结束后，按照调用顺序的相反顺序逐个执行
//即使函数发生严重错误也会执行
//常用于：资源清理，文件关闭，解锁以及记录时间等操作
//
/*
func main() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)  //i作为参数传递进来
	}
}

//结果 ：2 1 0
*/
/*
func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i) //这里的i是闭包 作为引用引用局部变量"i"
		}() //调用和这个匿名函数  类似以前写的  A()32,p
	}
}
//结果：
// 3 3 3
*/

//main函数结束的时候才执行的defer语句
//重点： go语言中没有try catch
//这里的defer  相当于 finally
//因为go语言中没有异常机制，但是有panic/recover模式处理错误
//Painc可以在任何地方引发，但recover只有在defer调用的函数中有效

//05
//painc  recover
//例子1
/*
func main() {
	A()
	B()
	C()
}

func A() {
	fmt.Println("Func in A")
}

func B() {
	fmt.Println("Func in B start")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //!  少了个（）  ： expression in go/defer must be function call
	panic(55) //自己设置的panic异常
	fmt.Println("Func in B end")
}

func C() {
	fmt.Println("Func in C")
}
*/
/*
结果：
Func in A
Func in B start
55
Func in C
*/
//结论：Func in B end 没有输出  但是  Func in C 却成功输出了

//例子2
/*
func main() {
	defer func() { //必须先声明defer，否则不能捕获panic异常
		fmt.Println("ma")
		if err := recover(); err != nil {  //recover 可以使进入panic状态的函数拉回来
			fmt.Println(err) //这里的err其实是painc传入的内容，55
			fmt.Println("in err")
		}
		fmt.Println("mb")
	}()
	f()
}

func f() {
	fmt.Println("fa")
	panic(55)
	fmt.Println("fb")
	fmt.Println("fc")
}
*/
/*
结果：
fa
ma
55
in err
mb
*/
//说明：当出现painc后  没有执行fb  fc  ！！！这里和实例1有点出入 有问题！！！

//作业  解释输出结果？？
//自我分析
/*
func main() {
	var fs = [4]func(){} //定义一个指向4个函数的变量

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		defer func() { fmt.Println("defer_closure i= ", i) }()
		fs[i] = func() { fmt.Println("closure i=", i) }
	}

	for _, f := range fs { //遍历这4个函数  对象是fs
		f() // 键值对中  f带标的是值，既是：[4]func()
	}
}
*/
//结果
/*
closure i= 4
closure i= 4
closure i= 4
closure i= 4
defer_closure i=  4
defer i= 3
defer_closure i=  4
defer i= 2
defer_closure i=  4
defer i= 1
defer_closure i=  4
defer i= 0
*/
//我的结论：这什么鬼  一点都看不懂

//老师解析：
func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i) //这里的i是一个参数，应该遵循正常的一个规则 值拷贝的传递
		//defer定义的时候已经获得了这个值拷贝  所以是 0  1  2  3  因为是defer  所以  反过来打印
		defer func() { fmt.Println("defer_closure i= ", i) }() //闭包思想
		fs[i] = func() { fmt.Println("closure i=", i) }        //这个匿名函数 自身并没有定义i，但是他用到了i  这个i是从外部夺过来的”引用地址“
		//我们首先把这些匿名函数存到了->fs[i]：func类型slice当中
	}

	for _, f := range fs { //for循环对func进行调用  输出i的值
		f()
	}
}

//结果
/*
closure i= 4  //这里打印的是func的值
closure i= 4
closure i= 4
closure i= 4
defer_closure i=  4  //defer  后进先出
defer i= 3
defer_closure i=  4
defer i= 2
defer_closure i=  4
defer i= 1
defer_closure i=  4
defer i= 0
*/
