package goroutine

import (
	"fmt"
	"runtime"
)

func Go1() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched() // 暂停协程调度执行，优选其他协程运行
		fmt.Println("hello")
	}
}
