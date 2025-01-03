package main

import "fmt"

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
}

func main() {
	//testPrint()

	//testVar()
}
