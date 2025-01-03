package main

import (
	"fmt"
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

func main() {
	//testPrint()
	//testVar()
	//testConst()
	//testExp()
	//testIfElse()
	testFor()

}
