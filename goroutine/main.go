package main

import (
	"fmt"
	g "king/goroutine/goroutine"
	"runtime"
)

func main() {
	//g.Go1()
	//g.Track()
	g.Runtime() // runtime

	fmt.Println("num cpu", runtime.NumCPU()) //cpu 数量
	fmt.Println("num goroutine", runtime.NumGoroutine())

}
