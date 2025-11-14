package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	num int64
)

func add(wg *sync.WaitGroup) {
	wg.Done()
	atomic.AddInt64(&num, 1)
}

// RunAtomic 原子操作
// atomic 适用于简单的同步操作
// atomic的性能要优于 sync.Mutex 同步互斥锁
func RunAtomic() {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go add(&wg)
	}
	wg.Wait()
	end := time.Since(start)
	fmt.Println("运行时间：", end)
	fmt.Println("结果：", num)
}
