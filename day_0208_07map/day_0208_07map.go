// day_0208_07map project day_0208_07map.go
package main //day_0208_07map

import "fmt"

//import "sort"

//键值对  key  map 存在
//必须支持  ==或！=比较运算类型
//Map查找比较性搜索快很多， 但比使用索引（slice / array）访问数据的类型慢100倍
//所以尽量使用slice或array
//Map 使用make创建
// 01 map
/*
func main() {
	//	var m map[int]string
	//	m = map[int]string{}
	//	fmt.Println(m)

	m := make(map[int]string) //比较简洁
	fmt.Println(m)
	//如何使用map
	m[1] = "OK"
	a := m[1]
	fmt.Println(a) //取出m[1]

	delete(m, 1) //m:map集合,1：代表的是键值对1
	fmt.Println("first")
	fmt.Println(a) //这样没有删掉

	a = m[1] //重新指向m[1]  上面没删掉的原因是已经把m[1]赋值给了a  a里面保存了m[1]的值
	fmt.Println("second")
	fmt.Println(a) //

	fmt.Println("---复杂的map--")
	var ma map[int]map[int]string
	ma = make(map[int]map[int]string)
	ma[1] = make(map[int]string) //需要对map内部的map初始化  //！ 少了[1]
	ma[1][1] = "OK"
	//	ma[2][1] = "III" //又没有初始化 报错
	va := ma[1][1]
	//	vb := ma[3][1]
	fmt.Println(va)
	//	fmt.Println(vb)
	//结论：不实用  每次实用内部的ma[几]都要初始化一次
}
*/

//迭代操作
//02
/*
func main() {
	//例子1
	//	for i, v := range slice {  //有两个变量  一个是i：   一个是v：值
	//		fmt.Println(i, v)
	//	}

	//	for k,v:=range map{
	//		map[k]
	//	}

	//e.g1:注意现象
	//	sm := make([]map[int]string, 4)
	//	for _, v := range sm {
	//		v = make(map[int]string, 1)
	//		v[1] = "OK"
	//		fmt.Println(v)
	//	}
	//	fmt.Println(sm)
	//现象：
	//map[1:OK]
	//map[1:OK]
	//map[1:OK]
	//map[1:OK]
	//[map[] map[] map[] map[]]
	//解释：v 是拷贝操作，你对v的任何操作都不会影响sm的本身

	//如何解决： 看下面代码
	sm := make([]map[int]string, 4)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
	//解决
	//现象
	//map[1:OK]
	//map[1:OK]
	//map[1:OK]
	//map[1:OK]
	//[map[1:OK] map[1:OK] map[1:OK] map[1:OK]]
}
*/

//03
//map间接排序

import "sort"

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h"} //!W  2, "b"  报错：missing key in map literal
	//	m := map[int]string{1: "a", 2: "b"}
	s := make([]int, len(m)) //s 是一个int型的切片
	i := 0
	for k, _ := range m { //map里面的数据是无序的  任何时候都是  我们通过外部的i 使其有序排列
		s[i] = k
		i++
	}
	//sort.Ints(s) //排序
	sort.Ints(s) //对int型 排序
	fmt.Println(s)

}

//作业：通过for range 将map的键值进行交换
//效果如下：
// m1 map[0:j 3:u 2:r 4:p]
// m2 map[u:3 p:4 j:0 2:r]

//自己的思考 (代码不完整 不会写）
/*
func main() {
	m1 := map[int]string{0: "j", 3: "u", 2: "r", 4: "p"}
	m2 := map[int]string{}  //！m2 := map[int]string{} 这样写是错的

	i, _ := 0, 0
	for _, v := range m1 {
		//if k ==i{
		m2[i] = v
		i++

		//		}
	}
	fmt.Println(m2)
}
*/
//老师的代码
/*
func main() {
	m1 := map[int]string{0: "j", 3: "u", 2: "r", 4: "p"}
	m2 := map[string]int{}

	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
*/
//17.2.8  1:20 止于09课函数function
