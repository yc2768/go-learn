package main

import (
	"fmt"
)

func main() {
	//字符拼接
	fmt.Println("Hello" + " Go!")

	var stockcode = 123
	var enddate = "2025.01.02"
	var url = "Code=%d&enddate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)

}
