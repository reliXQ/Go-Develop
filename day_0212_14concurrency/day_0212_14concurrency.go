// day_0212_14concurrency project day_0212_14concurrency.go
package main //day_0212_14concurrency

import (
	"fmt"
	//"time"
)

//go 高并发（concurrency）
//就是超级线程池  每个实例4~5kb的栈内存占用和由于实现机制而大幅减少的
//创建和销毁开销，是制造go号称的高并发的跟本原因
//并发不是并行 并发是由于切换时间片来实现“同时”运行，而“并行”则是通过
//计算机的多核CPU实现的多线程的运行，但go可以设置使用核数，以发挥多核
//计算机的能力

//Channel
//通道 共享内存
//make 创建  close关闭
//可以通过for range来迭代不断造作channel
//设置单向或双向通道
//设置缓存大小，未被填满前不会发生阻塞

//1
//func main() {
//	go Go()
//	time.Sleep(1 * time.Second) //time.Second
//}

//func Go() {
//	fmt.Println("Go Go Go!!!")
//}

//2
//channel
//通过通信实现内存共享，而不是共享内存来通信
/*
func main() {
	c := make(chan bool) //chan + bool:所存储的类型   //无缓存的是“同步的”   异步写法： c:=make(chan bool,10 ) //缓存为10
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true //存的是一个布尔类型   //这里是吧true 放入管道里了
	}()
	<-c //阻塞 直到true放入

}
//管道 有缓存则是异步的，不管你读不读我都继续执行，，无缓存是同步阻塞的
*/

//03
//for range  //等待程序退出
/*
func main() {
	go func() {
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c) //关闭通道
	}()
	for v := range c { //对chan 进行迭代操作  直到close（c）关闭通道
		fmt.Println(v)
	}
}
*/
/*
//此方法用于迭代 不断操作channel
for v := range c {
	fmt.Println(v)
}
*/

//04
//事例1
/*
import "time"

var c int

func counter() int {
	c++
	return c
}

func main() {
	a := 100

	go func(x, y int) {
		time.Sleep(time.Second) //让 goroutine 在main逻辑之后运行
		fmt.Println("go:", x, y)
	}(a, counter()) //传参
	a += 100
	fmt.Println("main:", a, counter())

	time.Sleep(time.Second * 3)
}
*/
//结果：
//main: 200 2
//go: 100 1

//05
//事例2
//Wait 进程退出时 不会等待并发任务结束，可用通道（channel）阻塞，然后退出
/*
import "time"

func main() {
	exit := make(chan struct{}) //创建通道，因为仅仅是通知，数据没有实际意义

	go func() {
		time.Sleep(time.Second)
		fmt.Println("goroutine done.")

		close(exit) //关闭通道  //退出  //注销掉这个发生： all goroutines are asleep - deadlock!
	}()

	fmt.Println("main...")
	<-exit //  通道阻塞  发出退出信号
	fmt.Println("main exit.")
}

//除了关闭通道外，写入数据也可以解除阻塞。
*/

//有缓冲是异步的  无缓存是同步的

//06-1
/*
func main() {
	for i := 0; i < 10; i++ {
		go Go(i)
	}
}

func Go(index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)
}

//结果：main程序没有等待go的程序所以没有打出任何东西，程序就结束了
*/

//06-2
/*
//import "time"
//有问题
func main() {
	c := make(chan bool) //管道
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	<-c
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	if index == 9 {
		c <- true
	}
}
*/

//06-3
/*
//输出10此 且10此顺序不是按照常规依次运行的
func main() {
	c := make(chan bool) //管道
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 100000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}
*/

//06-4

//提高效率
/*
import "runtime"

//高效
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	//	<-c
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	c <- true
}
*/
//06-第二种解决办法
/*
import "runtime"
import "sync" //可以添加任务组  任务组可以不断添加任务数 然后你没完成一个任务时标记down 则它减到0时任务就完成了
//sync  中 WaitGroup  同步包
//高效
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10) //组为10  如果不设置组的个数程序仍然输出不对
	//	c := make(chan bool)
	for i := 0; i < 10; i++ {
		go Go(&wg, i)
	}
	wg.Wait() //等待
}

func Go(wg *sync.WaitGroup, index int) { //传进来指针
	a := 1
	for i := 0; i < 1000000; i++ {
		a += i
	}
	fmt.Println(index, a)

	wg.Done() //每完成一次 减一
}
*/

//07 处理多个结构发送与接收的问题 Select
//可同时处理多个channel的发送与接收
//同时有多个可用的channel时按随机顺序处理
//可用空的select来阻塞main函数
//可设置超时

//import(
//	"runtime"
//)
/*
func main() {
	//这两个通道传值
	c1, c2 := make(chan int), make(chan string)
	//通信通道
	o := make(chan bool, 2)
	go func() { //匿名函数
		for {
			select {
			case v, ok := <-c1: //取出c1的值
				if !ok {
					o <- true
					break //c1被关闭的话 这里break
				}
				fmt.Println("c1", v)
			case v, ok := <-c2:
				if !ok {
					o <- true
					break
				}
				fmt.Println("c2", v)
			}
		}
	}()

	//传值
	c1 <- 1
	c2 <- "hi"
	c1 <- 2
	c2 <- "hello"

	close(c1) //只需要关闭一个即可
	//	close(c2)

	<-o
}
*/

//08
/*
//select 作为发送者的应用

func main() {
	c := make(chan int)
	go func() {
		for v := range c { //循环读取c中的值
			fmt.Println(v)
		}
	}()

	for i := 0; i < 10; i++ {
		select { //随机发送一个值
		case c <- 0:
		case c <- 1:
		case c <- 2:
		case c <- 3:
		}
	}
}
*/
//select{} 空select用来阻塞 //这里的应用：GUI中 嵌入事件的一个循环main不能退出 卡死  程序中发送一些关闭信号来让程序退出

//超时
/*
import "time"

func main() {
	c := make(chan bool)
	select {
	case v := <-c:
		fmt.Println(v)
	case <-time.After(3 * time.Second): //expecting := or = or : or comma  !W 少了个":"
		fmt.Println("3 Times Out")
	}
}
*/

//作业; 两个chan  来回发发送数据
/*
func PingPang(c chan string) {
	i := 0
	for {
		fmt.Println(<-c) //阻塞
		c <- fmt.Sprintf("From PingPang:hi #%d", i)
		i++
	}
}

func PangPing(c chan string) {

	for j := 0; j < 10; j++ {
		c <- fmt.Sprintf("From PangPing:hii #%d", j)
		fmt.Println(<-c)
		j++
	}
}
func main() {
	c := make(chan string)
	go PingPang(c)
	go PangPing(c)
}
*/
//var c chan string

func PingPang(c chan string) {
	i := 0
	for {
		fmt.Println(<-c) //阻塞
		c <- fmt.Sprintf("From PingPang:hi #%d", i)
		i++
	}
}

func PangPing(c chan string) {

	for j := 0; j < 10; j++ {
		c <- fmt.Sprintf("From PangPing:hii #%d", j)
		fmt.Println(<-c)
		j++
	}
}
func main() {
	c := make(chan string)
	go PingPang(c)
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From PangPing:hii #%d", i)
		fmt.Println(<-c)
		i++
	}
	A()
	A()
	A()
}

//闭包思想
//闭包是由于函数及其相关引用环境组合而成的实体（既：闭包=函数+引用环境）
//闭包：拥有多个变量和绑定了这些变量的环境表达式（通常是一个函数）
//当函数A()的内部函数B()被函数A()外的一个变量引用的时候，就创建了一个闭包。

//错误事例：这样不是闭包
func A() func() {
	var i = 1
	a := func() { //匿名函数可以赋值给一个变量的
		i++
		fmt.Println("i:", i)
	}
	return a
}

//正确事例
//func A() {

//}
