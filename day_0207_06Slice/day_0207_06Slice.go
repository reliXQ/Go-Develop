// day_0207_06Slice project day_0207_06Slice.go
package main //day_0207_06Slice

//Slice
//len（）获取元素个数  cap() 获取容量
//一般使用make()创建
//不是用new创建slice  因为new创建的是指针  本身slice就是指针  如果用new创建则会创建出指针的指针我去
import "fmt"

// 01 切片定义方法1
/*
func main() {
	var s1 []int //这是“切片”  如果是数组的话【】里面应该有数字或...
	fmt.Println(s1)

	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} //变量 ：=[]int{}
	s2 := a[5:10]                               //切片 s1  //5:10指包括起始索引不包括终止索引
	fmt.Println(s2)
	fmt.Println("------")
	s3 := a[5:]
	fmt.Println(s3)
}
*/

//02 切片定义方法2  make
/*
func main() {
	s1 := make([]int, 3, 10)  //参数1:[]int数组 参数2:3代表初始容量为3，  参数3:10 10段连续内存，不超过10个参数的内存都是连续的
	s2 := make([]int, 10, 20) //每次如果长度不够了，内存地址块长度会翻倍
	fmt.Println(s1)
	fmt.Println(len(s1), cap(s1))
	fmt.Println(s2)
	fmt.Println(len(s2), cap(s2)) //长度为10  容量为20
}

*/

//slice
/*
func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	sb := a[3:5]
	//fmt.Println(sa)
	fmt.Println(string(sa))
	fmt.Println(string(sb))
	//Reslice 的切片是以被切片的数据为准
	sc := a[2:6]
	sd := sc[:3]
	fmt.Println(string(sc))
	fmt.Println(string(sd))
	//索引不可以超过被slice的切片的容量cap（）值
	//索引越界会引发错误

	//append 在slice尾部追加元素
	//如果超过追加的slice的容量，则将会重新分配数并拷贝原始数据
	//Copy

	s1 := make([]int, 3, 6) //数组为3 容量为6
	fmt.Printf("%p len=%v cap=%v\n", s1, len(s1), cap(s1))
	fmt.Println("-------")
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	fmt.Println("-------")
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Printf("%v %p len=%v cap=%v\n", s1, s1, len(s1), cap(s1))

	//	 现象：
	//		0xc04203bd70
	//	-------
	//	[0 0 0 1 2 3] 0xc04203bd70
	//	-------
	//	[0 0 0 1 2 3 1 2 3 4 5] 0xc042036180

	//结论：当不超过元素容量的时候 直接追加
	//如果数据超过了容量的大小，则我们的slice会重新分配一个slice
	//容量 翻倍

	//修改其中一个slice的值，则另外一个slice会改变
	b := []int{1, 2, 3, 4, 5, 6}
	sb1 := b[2:5] //公用部分为b[2]
	sb2 := b[1:3]
	fmt.Printf("sb %v %v \n", sb1, sb2)
	sb1[0] = 9
	fmt.Printf("sb %v %v \n", sb1, sb2)
	//如果append sb2中的数据  如果超过容量值 则会重新分配内存
	sb2 = append(sb2, 7, 7, 7, 7, 7)
	sb1[0] = 8
	fmt.Printf("append over :sb %v %v \n", sb1, sb2)
	//结论：超过容量值  则sb2指向的地址已经改变了
}
*/

//copy
func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{9, 9, 9}
	//	copy(s1, s2)
	//	copy(s2, s1)
	//	copy(s1[2:6], s2[0:2])  //2个数拷贝到4个容量中  其余不修改
	copy(s1[:1], s2[0:2]) //总体来说 多数服从少数
	fmt.Println(s1)
	fmt.Println(s2)
}
