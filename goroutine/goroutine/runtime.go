package goroutine

import (
	"fmt"
	"runtime"
)

func Track() {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	fmt.Printf("goroutine\n %s", buf[:n])
}

func Runtime() {
	// 内存分配信息
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Printf("%#v\n", mem)
	// 获取运行时 当前函数的相关信息
	pc, _, _, _ := runtime.Caller(0) // 获得当前函数的pc地址
	info := runtime.FuncForPC(pc)
	if info != nil {
		fmt.Printf("func name \n%s\n", info.Name())
		file, line := info.FileLine(pc)
		fmt.Printf("func file file %s\n", file)
		fmt.Printf("func file line %d\n", line)
	}
}
