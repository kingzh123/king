package goroutine

import "fmt"

func CreateChanel() {
	// 声明通道
	var a chan int
	var aa chan bool
	var aaa chan []int
	// 创建通道
	b := make(chan int)
	fmt.Println(a)
	fmt.Println(aa)
	fmt.Println(aaa)
	fmt.Println(b)
}
