package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func testPrint() {
	//字符拼接
	fmt.Println("Hello" + " Go!")

	var stockcode = 123
	var enddate = "2025.01.02"
	var url = "Code=%d&enddate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
}

func testVar() {
	var a string = "Go"
	fmt.Println(a)
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)
	var e bool
	fmt.Println(e)

	var a1 *int
	a1 = new(int)
	*a1 = 10
	*a1 = 100
	var b1 = 20
	a1 = &b1
	fmt.Println(&b1, a1, *a1, &a1)

	var a2 []int
	fmt.Println(a2)

	var a3 map[string]string
	fmt.Println(a3)

	var a4 chan int
	fmt.Println(a4)

	var a5 func(int) int
	fmt.Println(a5)

	var a6 error
	fmt.Println(a6)

	var a7 = 8
	fmt.Println(a7)

	a8 := false
	fmt.Println(a8)

	var b2, b3, b4 int = 1, 2, 3
	fmt.Println(b2, b3, b4)
	c1, c2, c3 := 4, 5, 6
	fmt.Println(c1, c2, c3)

	var (
		d1 int
		d2 bool
	)
	fmt.Println(d1, d2)

	var e1 = 2
	var e2 = &e1
	fmt.Println(e2, *e2, &e2)
}

func testConst() {
	const LENGTH int = 100
	fmt.Println(LENGTH)

	const (
		Unknown = 0
		Femal   = 1
		Male    = 2
	)
	fmt.Println(Unknown, Femal, Male)

	const (
		A = iota
		B
		C
	)
	fmt.Println(A, B, C)

	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

func testExp() {
	//var a int = 100
	//var b int = 200
	//var c int = a + b
	//fmt.Printf("%d\n", c)
	//fmt.Printf("%d\n", a-b)
	//fmt.Printf("%d\n", a*b)
	//fmt.Printf("%d\n", a/b)
	//fmt.Printf("%d\n", a%b)
	//a++
	//fmt.Printf("%d\n", a)
	//b--
	//fmt.Printf("%d\n", b)

	//var a int = 20
	//var b int = 21
	//fmt.Println(a == b)
	//fmt.Println(a < b)
	//fmt.Println(a > b)
	//fmt.Println(a <= b)
	//fmt.Println(a >= b)

	//var a bool = true
	//var b bool = false
	//fmt.Println(a && b, a || b, !a)

	var a uint = 60
	var b uint = 12
	var c uint
	c = a | b
	fmt.Println(c)
	c = a & b
	fmt.Println(c)
	c = a ^ b
	fmt.Println(c)
	c = a >> 2
	fmt.Println(c)
	c = a << 2
	fmt.Println(c)

}

func testIfElse() {
	//var score = 99
	//if score > 90 {
	//	fmt.Println("通过")
	//} else {
	//	fmt.Println("不通过")
	//}

	//var grade string = "B"
	//var score int = 90
	//switch score {
	//case 90:
	//	grade = "A"
	//case 80:
	//	grade = "B"
	//case 50, 60, 70:
	//	grade = "C"
	//default:
	//	grade = "D"
	//}
	//switch {
	//case grade == "A":
	//	fmt.Println("优秀")
	//case grade == "B":
	//	fmt.Println("良好")
	//case grade == "C":
	//	fmt.Println("及格")
	//default:
	//	fmt.Println("不及格")
	//}

	//var x interface{}
	//switch x.(type) {
	//case nil:
	//	fmt.Println("nil")
	//case int:
	//	fmt.Println("int")
	//case string:
	//	fmt.Println("string")
	//default:
	//	fmt.Println("未知类型")
	//}

	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "来自通道1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自通道2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

	//ch1 := make(chan string)
	//ch2 := make(chan string)
	//go func() {
	//	for {
	//		ch1 <- "one"
	//	}
	//}()
	//go func() {
	//	for {
	//		ch2 <- "two"
	//	}
	//}()
	//for {
	//	select {
	//	case msg1 := <-ch1:
	//		fmt.Println("msg1:" + msg1)
	//	case msg2 := <-ch2:
	//		fmt.Println("msg2:" + msg2)
	//	default:
	//		fmt.Println("no message received")
	//	}
	//}

}

func testFor() {
	//sum := 0
	//for i := 0; i < 20; i++ {
	//	sum += i
	//}
	//fmt.Println(sum)

	//sum := 0
	//for sum < 100 {
	//	sum += sum
	//}
	//fmt.Println(sum)

	//strings := []string{"apples", "oranges", "bananas"}
	//for i, s := range strings {
	//	fmt.Println(i, s)
	//}

	map1 := make(map[string]float64)
	map1["apple"] = 10000
	map1["huawei"] = 5000
	fmt.Println(map1)
	for key, value := range map1 {
		fmt.Println(key, value)
	}

}

func swap(x, y string) (string, string) {
	return y, x
}

func getSeqence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

/* 声明全局变量 */
var a int = 20

/* 函数定义-两数相加 */
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)

	return a + b
}

func main() {
	//testPrint()
	//testVar()
	//testConst()
	//testExp()
	//testIfElse()
	//testFor()
	//y, x := swap("nihao", "Go")
	//fmt.Println(y, x)

	//nextNumber := getSeqence()
	//fmt.Println(nextNumber())
	//fmt.Println(nextNumber())
	//nextNumber1 := getSeqence()
	//fmt.Println(nextNumber1())
	//fmt.Println(nextNumber1())

	//add := func(a, b int) int {
	//	return a + b
	//}
	//
	//multiply := func(a, b int) int {
	//	return a * b
	//}
	//calculate := func(operation func(int, int) int, x, y int) int {
	//	return operation(x, y)
	//}
	//fmt.Println(calculate(add, 1, 2))
	//fmt.Println(calculate(multiply, 2, 2))

	//var c1 Circle
	//c1 = Circle{radius: 5}
	//fmt.Println(c1.area())

	/* main 函数中声明局部变量 */
	//var a int = 10
	//var b int = 20
	//var c int = 0
	//fmt.Printf("main()函数中 a = %d\n", a)
	//c = sum(a, b)
	//fmt.Printf("main()函数中 c = %d\n", c)

	//var numbers = [4]int{1, 2, 3, 4}
	//numbers1 := [4]int{4, 5, 6, 7}
	//fmt.Println(numbers)
	//fmt.Println(numbers1)
	//
	//var blance = [...]float32{1.0, 2.0}
	//blance1 := [...]float32{3.0, 4.0}
	//balance2 := [5]float32{1: 2.0, 3: 7.0}
	//fmt.Println(blance, blance1, balance2)
	//for i := 0; i < len(numbers); i++ {
	//	fmt.Println(numbers[i])
	//}

	//指针
	/*var a int = 10
	var p *int = &a
	fmt.Println(*p, p, &a)*/

	//指针数组
	//const MAX int = 3
	//arr := [3]int{10, 100, 200}
	//var i int
	//var ptr [MAX]*int
	//for i = 0; i < MAX; i++ {
	//	ptr[i] = &arr[i]
	//}
	//for i = 0; i < MAX; i++ {
	//	fmt.Println(ptr[i], *ptr[i])
	//}

	//指向指针的指针变量
	//var a int = 20
	//var ptr *int = &a
	//var pptr **int = &ptr
	//fmt.Println(a)
	//fmt.Println(*ptr)
	//fmt.Println(**pptr)

	//结构体
	//type Books struct {
	//	title   string
	//	author  string
	//	subject string
	//}
	//book := Books{"红楼梦", "曹雪芹", "四大名著"}
	//book1 := Books{title: "红楼梦", author: "曹雪芹", subject: "四大名著"}
	//book2 := Books{title: "红楼梦", author: "曹雪芹"}
	//fmt.Println(book, book1, book2, book2.title)
	////结构体指针
	//var ptr *Books
	//ptr = &book1
	//fmt.Println((*ptr).subject)

	//切片
	/*var numbers = make([]int, 3, 5)
	numbers1 := make([]int, 3, 5)
	numbers[0] = 1
	numbers2 := []int{4, 5, 6, 7}
	var numbers3 []int
	fmt.Println(numbers, numbers1, numbers1[1:], numbers2, len(numbers1), cap(numbers1), numbers3)
	numbers1 = append(numbers1, 6, 8, 10)
	fmt.Println(numbers1)
	numbersCopy := make([]int, len(numbers1)*2, cap(numbers1)*2)
	copy(numbersCopy, numbers1)
	fmt.Println(numbersCopy)*/

	//遍历切片
	//number := []int{1, 2, 3}
	//for i, v := range number {
	//	fmt.Println(i, v)
	//}
	//遍历字符串
	//for i, v := range "hello" {
	//	fmt.Printf("i=%d, v=%s\n", i, string(v))
	//}
	//map
	//map1 := make(map[string]int)
	//map1["apple"] = 10000
	//map1["banana"] = 5000
	//for k, v := range map1 {
	//	fmt.Println(k, v)
	//}
	//通道
	//ch1 := make(chan int, 2)
	//ch1 <- 2
	//ch1 <- 3
	//close(ch1)
	//for v := range ch1 {
	//	fmt.Println(v)
	//}

	//创建空map
	//m1 := make(map[string]int)
	//fmt.Println(m1)
	////创建一个初始容量为 2 的 Map
	//m2 := make(map[string]int, 2)
	//fmt.Println(m2)
	//m3 := map[string]int{"apple": 10, "orange": 20}
	//v1 := m3["apple"]
	//m3["apple"] = 56
	//delete(m3, "orange")
	//fmt.Println(m3, v1, len(m3))

	//类型转换
	//var a int = 10
	//var b float64 = float64(a)
	//fmt.Println(b)

	var str string = "10"
	var num int
	num, _ = strconv.Atoi(str)
	fmt.Println(num)

}
